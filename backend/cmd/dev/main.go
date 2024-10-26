package main

import (
	"backend/pkg/settings"
	"backend/svc/pkg/handler"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	r := gin.Default()
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
	fmt.Println(dbUrl)
	// Health Check
	h := handler.NewHealthHandler()
	r.GET("/health", h.Handle)
	r.Run()
}
