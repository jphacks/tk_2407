package main

import (
	"backend/pkg/settings"
	"backend/svc/pkg/handler"
	"backend/svc/pkg/middleware"
	"backend/svc/pkg/query"
	"backend/svc/pkg/usecase"
	"errors"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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

	if conf.Application.Server.OnProduction {
		log.Println("Running migration...")
		m, err := migrate.New("file:///app/migrations", dbUrl)
		if err != nil {
			log.Fatalf("failed to create migration: %v", err)
			return
		}
		if err := m.Up(); err != nil {
			if !errors.Is(err, migrate.ErrNoChange) {
				log.Fatalf("failed to migrate: %v", err)
				return
			}
			log.Println("No migration needed.")
		}
		log.Println("Migration done.")
	}

	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database, err: %v", err)
	}

	q := query.Use(db, nil)

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

// Implement APIのルーティングをするところ
func Implement(rg *gin.RouterGroup, q *query.Query) error {
	rg.GET("/health", handler.NewHealthHandler().Handle)

	userUsecase := usecase.NewUserUsecase(q)
	userHandler := handler.NewUserHandler(userUsecase)

	rg.GET("/user/:userId", userHandler.GetUser())

	return nil
}
