package handler

import (
	"backend/svc/pkg/openapi"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HealthHandler struct {
}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) Handle(c *gin.Context) {
	c.JSON(http.StatusOK, &openapi.HealthRes{
		Message: "healthy",
	})
}
