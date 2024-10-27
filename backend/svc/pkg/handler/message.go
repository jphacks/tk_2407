package handler

import (
	"backend/svc/pkg/openapi"
	"backend/svc/pkg/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

// MessageHandler message handler
type MessageHandler struct {
	mu *usecase.MessageUsecase
}

// NewMessageHandler constructor
func NewMessageHandler(mu *usecase.MessageUsecase) *MessageHandler {
	return &MessageHandler{
		mu: mu,
	}
}

// Create 寄せ書き作成
func (mh MessageHandler) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var body openapi.PostApiV1MessageJSONRequestBody
		if err := c.Bind(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
			c.Abort()
			return
		}

		message, err := mh.mu.Create(body.SpotId, body.UserId, body.Content, body.PhotoUrl)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "failed to create message"})
			c.Abort()
			return
		}

		res := openapi.SuccessMessageCreateRes{
			Content:   *message.Content,
			MessageId: message.ID,
			PhotoUrl:  *message.PhotoURL,
			SpotId:    message.SpotID,
			UserId:    message.UserID,
		}

		c.JSON(http.StatusCreated, res)
	}
}
