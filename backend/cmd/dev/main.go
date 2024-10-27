package main

import (
	"backend/pkg/google_maps"
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

	// parse json
	data, err := os.ReadFile("/app/migrations/config.json")
	if err != nil {
		return fmt.Errorf("failed to read config.json: %w", err)
	}
	type MigrationConfig struct {
		Version *uint `json:"version"`
		Force   *uint `json:"force"`
	}
	var configMigration MigrationConfig
	if err := json.Unmarshal(data, &configMigration); err != nil {
		return fmt.Errorf("failed to unmarshal config.json: %w", err)
	}

	if configMigration.Version == nil {
		return errors.New("version key is missing")
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

	if configMigration.Force != nil {
		log.Printf("Forcing version: %v\n", *configMigration.Force)
		if err := m.Force(int(*configMigration.Force)); err != nil {
			return fmt.Errorf("failed to force migration: %w", err)
		}
		log.Println("Force Migration done.")
		return nil
	}
	log.Printf("Migrating to version: %v\n", *configMigration.Version)
	if err := m.Migrate(*configMigration.Version); err != nil {
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
	messageUsecase := usecase.NewMessageUsecase(q)
	messageHandler := handler.NewMessageHandler(messageUsecase)

	rg.GET("/user/:userId", userHandler.GetUser())
	rg.POST("/signup", authHandler.Signup())
	rg.POST("/login", authHandler.Login())
	rg.POST("/message", messageHandler.Create())

	gmClient, err := google_maps.NewGoogleMaps()
	if err != nil {
		log.Fatalf("failed to create google maps client: %v", err)
	}
	googleMap := handler.NewGoogleMapHandler(q, *gmClient)
	rg.GET("/spots", googleMap.GetApiV1Spots)
	rg.GET("/photo/:photoRef", googleMap.GetPhotos)
	return nil
}
