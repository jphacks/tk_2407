package handler

import (
	"backend/svc/pkg/openapi"
	"backend/svc/pkg/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AuthHandler auth handler
type AuthHandler struct {
	au *usecase.AuthUsecase
}

// NewAuthHandler constructor
func NewAuthHandler(au *usecase.AuthUsecase) *AuthHandler {
	return &AuthHandler{
		au: au,
	}
}

// Signup signup
func (ah AuthHandler) Signup() gin.HandlerFunc {
	return func(c *gin.Context) {
		var body openapi.PostApiV1SignupJSONRequestBody
		if err := c.Bind(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "email or username or password is necessary"})
			c.Abort()
			return
		}

		res, tokenString, err := ah.au.Signup(body.Email, body.Username, body.Password)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}

		c.SetCookie("token", tokenString, 3600*24, "/", "", false, true)

		c.JSON(http.StatusCreated, res)
	}
}

// Login login
func (ah AuthHandler) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var body openapi.PostApiV1LoginJSONRequestBody
		if err := c.Bind(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "email or password is necessary"})
			c.Abort()
			return
		}

		res, tokenString, err := ah.au.Login(body.Email, body.Password)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}

		c.SetCookie("token", tokenString, 3600*24, "/", "", false, true)

		c.JSON(http.StatusCreated, res)
	}
}
