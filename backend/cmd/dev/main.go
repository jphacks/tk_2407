package main

import (
	"backend/pkg/settings"
	"backend/svc/pkg/handler"
	"backend/svc/pkg/middleware"
	"backend/svc/pkg/query"
	"backend/svc/pkg/usecase"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"gorm.io/driver/postgres"
)

func main() {
	engine := gin.Default()
	conf := settings.Get()
	var dbUrl string
	switch conf.Infrastructure.Postgres.Protocol {
	case "tcp":
		dbUrl = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
			conf.Infrastructure.Postgres.User,
			conf.Infrastructure.Postgres.Password,
			conf.Infrastructure.Postgres.Host,
			conf.Infrastructure.Postgres.Port,
			conf.Infrastructure.Postgres.DBName)
	case "unix":
		dbUrl = fmt.Sprintf("postgres://%s:%s@/%s?host=%s",
			conf.Infrastructure.Postgres.User,
			conf.Infrastructure.Postgres.Password,
			conf.Infrastructure.Postgres.DBName,
			conf.Infrastructure.Postgres.UnixSocket)
	default:
		log.Fatalf("invalid protocol: %s", conf.Infrastructure.Postgres.Protocol)
	}

	// ping db
	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatalf("failed to open db: %v", err)
		return
	}
	if err := db.Ping(); err != nil {
		log.Fatalf("failed to ping db: %v", err)
		return
	}

	if conf.Application.Server.OnProduction {
		if err := runMigration(dbUrl); err != nil {
			log.Fatalf("error during migration: %v", err)
			return
		}
	}

	gormDB, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database, err: %v", err)
	}

	q := query.Use(gormDB, nil)

	// Implement Application API
	apiV1 := engine.Group("/api/v1")
	middleware.NewCORS().ConfigureCORS(apiV1)
	if err := Implement(apiV1, q); err != nil {
		log.Fatalf("Failed to start server... %v", err)
		return
	}

	if err := engine.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server... %v", err)
		return
	}
}

func runMigration(dbUrl string) error {
	log.Println("Running migration...")

	//read json from /app/migrations/.current.json
	if _, err := os.Stat("/app/migrations/.config.json"); err == nil {
		log.Println("No migration run")
		return nil
	}
	// parse json
	data, err := os.ReadFile("/app/migrations/.current.json")
	if err != nil {
		return fmt.Errorf("failed to read .current.json: %w", err)
	}
	var current map[string]interface{}
	if err := json.Unmarshal(data, &current); err != nil {
		return fmt.Errorf("failed to unmarshal .current.json: %w", err)
	}

	ver, ok := current["version"]
	if !ok {
		return errors.New("version key not found in .current.json")
	}
	verUInt, ok := ver.(uint)
	if !ok {
		return fmt.Errorf("version key is not uint: %+v", ver)
	}

	var forceVer *uint
	if force, ok := current["force"]; ok {
		forceV, ok := force.(uint)
		if !ok {
			return fmt.Errorf("force key is not uint: %+v", force)
		}
		forceVer = &forceV
	}

	m, err := migrate.New("file:///app/migrations", dbUrl)
	if err != nil {
		return fmt.Errorf("failed to create migration: %w", err)
	}
	version, b, err := m.Version()
	if err != nil {
		return fmt.Errorf("failed to get version: %w", err)
	}
	log.Printf("Current version: %v\n, dirty: %v\n", version, b)

	if forceVer != nil {
		log.Printf("Forcing version: %v\n", *forceVer)
		if err := m.Force(int(*forceVer)); err != nil {
			return fmt.Errorf("failed to force migration: %w", err)
		}
		log.Println("Force Migration done.")
		return nil
	}
	log.Printf("Migrating to version: %v\n", verUInt)
	if err := m.Migrate(verUInt); err != nil {
		if !errors.Is(err, migrate.ErrNoChange) {
			return fmt.Errorf("failed to migrate: %w", err)
		}
	}
	log.Println("Migration done.")
	return nil
}

// Implement APIのルーティングをするところ
func Implement(rg *gin.RouterGroup, q *query.Query) error {
	rg.GET("/health", handler.NewHealthHandler().Handle)

	userUsecase := usecase.NewUserUsecase(q)
	userHandler := handler.NewUserHandler(userUsecase)
	authUsecase := usecase.NewAuthUsecase(q)
	authHandler := handler.NewAuthHandler(authUsecase)

	rg.GET("/user/:userId", userHandler.GetUser())
	rg.POST("/signup", authHandler.Signup())
	rg.POST("/login", authHandler.Login())

	return nil
}
