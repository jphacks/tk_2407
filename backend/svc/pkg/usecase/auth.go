package usecase

import (
	"backend/pkg/cipher"
	"backend/pkg/jwt"
	"backend/pkg/settings"
	"backend/svc/pkg/domain"
	"backend/svc/pkg/openapi"
	"backend/svc/pkg/query"
	"log"
	"time"

	"github.com/oklog/ulid/v2"
)

// AuthUsecase auth usecase
type AuthUsecase struct {
	q *query.Query
}

// NewAuthUsecase constructor
func NewAuthUsecase(q *query.Query) *AuthUsecase {
	return &AuthUsecase{
		q: q,
	}
}

// Signup signup
func (au AuthUsecase) Signup(email string, username string, password string) (*openapi.SuccessSignupRes, string, error) {
	now := time.Now()
	userID := ulid.Make()
	passwordHash, err := cipher.GenHash(password)
	if err != nil {
		log.Println(err)
		return nil, "", err
	}
	if err := au.q.User.Create(&domain.User{
		ID:           userID.String(),
		Username:     username,
		Email:        email,
		PasswordHash: string(passwordHash),
		CreatedAt:    &now,
		UpdatedAt:    &now,
	}); err != nil {
		log.Println(err)
		return nil, "", err
	}

	config := settings.Get()
	duration := 24 * time.Hour // 一旦24H
	claims := jwt.CreateClaims(userID.String(), duration, username)
	tokenString, err := jwt.IssueJWT(claims, config.Service.Authentication.JwtSecret)
	if err != nil {
		log.Println(err)
		return nil, "", err
	}

	return &openapi.SuccessSignupRes{UserId: userID.String()}, tokenString, nil
}

// Login login
func (au AuthUsecase) Login(email string, password string) (*openapi.SuccessLoginRes, string, error) {
	user, err := au.q.User.Where(au.q.User.Email.Eq(email)).First()
	if err != nil {
		log.Println(err)
		return nil, "", err
	}
	if err := cipher.Compare(password, []byte(user.PasswordHash)); err != nil {
		log.Println(err)
		return nil, "", err
	}

	config := settings.Get()
	duration := 24 * time.Hour // 一旦24H
	claims := jwt.CreateClaims(user.ID, duration, user.Username)
	tokenString, err := jwt.IssueJWT(claims, config.Service.Authentication.JwtSecret)
	if err != nil {
		log.Println(err)
		return nil, "", err
	}

	return &openapi.SuccessLoginRes{UserId: user.ID}, tokenString, nil
}
