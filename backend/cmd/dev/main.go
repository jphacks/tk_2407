package main

import (
	"backend/pkg/settings"
	"backend/svc/pkg/handler"
	"backend/svc/pkg/middleware"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	conf := settings.Get()
	var dbUrl string
	switch conf.Infrastructure.Postgres.Protocol {
	case "tcp":
		dbUrl = fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable&TimeZone=Asia/Tokyo",
			conf.Infrastructure.Postgres.User,
			conf.Infrastructure.Postgres.Password,
			conf.Infrastructure.Postgres.Host,
			conf.Infrastructure.Postgres.Port,
			conf.Infrastructure.Postgres.DBName)
	case "unix":
		dbUrl = fmt.Sprintf("postgresql://%s:%s@unix:%s/%s?sslmode=disable&TimeZone=Asia/Tokyo",
			conf.Infrastructure.Postgres.User,
			conf.Infrastructure.Postgres.Password,
			conf.Infrastructure.Postgres.UnixSocket,
			conf.Infrastructure.Postgres.DBName)
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
			log.Fatalf("failed to migrate: %v", err)
			return
		}
	}

	// Implement Application API
	apiV1 := engine.Group("/api/v1")
	middleware.NewCORS().ConfigureCORS(apiV1)
	if err := Implement(apiV1); err != nil {
		log.Fatalf("Failed to start server... %v", err)
		return
	}

	if err := engine.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server... %v", err)
		return
	}
}

// Implement APIのルーティングをするところ
func Implement(rg *gin.RouterGroup) error {
	rg.GET("/health", handler.NewHealthHandler().Handle)
	return nil
}
