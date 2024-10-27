package usecase

import (
	"backend/svc/pkg/domain"
	"backend/svc/pkg/query"
)

// UserUsecase user usecase
type UserUsecase struct {
	q *query.Query
}

// NewUserUsecase constructor
func NewUserUsecase(q *query.Query) *UserUsecase {
	return &UserUsecase{
		q: q,
	}
}

// GetByID 単体取得
func (uu UserUsecase) GetByID(userID string) (*domain.User, error) {
	user, err := uu.q.User.Where(uu.q.User.ID.Eq(userID)).First()
	if err != nil {
		return nil, err
	}

	return user, nil
}
