package main

import (
	"backend/pkg/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Health Check
	h := handler.NewHealthHandler()
	r.GET("/health", h.Handle)
	r.Run()
}
