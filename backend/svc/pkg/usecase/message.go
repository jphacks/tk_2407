package usecase

import (
	"backend/svc/pkg/domain"
	"backend/svc/pkg/query"
	"log"
	"time"

	"github.com/oklog/ulid/v2"
)

// MessageUsecase message usecase
type MessageUsecase struct {
	q *query.Query
}

// NewMessageUsecase constructor
func NewMessageUsecase(q *query.Query) *MessageUsecase {
	return &MessageUsecase{
		q: q,
	}
}

// Create 寄せ書き作成
func (mu MessageUsecase) Create(spotID string, userID string, content string, photoURL string) (*domain.Message, error) {
	now := time.Now()
	messageID := ulid.Make()
	message := &domain.Message{
		ID:        messageID.String(),
		UserID:    userID,
		SpotID:    spotID,
		PhotoURL:  photoURL,
		Content:   content,
		CreatedAt: &now,
		UpdatedAt: &now,
	}
	if err := mu.q.Message.Create(message); err != nil {
		log.Println(err)
		return nil, err
	}

	return message, nil
}
