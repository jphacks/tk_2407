package handler

import (
	"backend/svc/pkg/openapi"
	"backend/svc/pkg/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserHandler user handler
type UserHandler struct {
	uu *usecase.UserUsecase
}

// NewUserHandler constructor
func NewUserHandler(uu *usecase.UserUsecase) *UserHandler {
	return &UserHandler{
		uu: uu,
	}
}

// GetUser 単体取得
func (uh UserHandler) GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.Param("userId")
		if len(userID) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "user id is necessary"})
			c.Abort()
			return
		}
		user, err := uh.uu.GetByID(userID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "failed to get user"})
			c.Abort()
			return
		}

		res := openapi.SuccessUserRes{
			Email:    user.Email,
			Username: user.Username,
		}

		c.JSON(http.StatusOK, res)
	}
}
