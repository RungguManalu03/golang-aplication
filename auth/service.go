package auth

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type Service interface {
	GenerateToken(UserID uuid.UUID) (string, error)
}

type jwtService struct {
}

var SECRET_KEY = []byte("4Jkl)!?.;asak7&^&%^$#HS(S&%RSYDL>LDJD*^)")

func NewService() *jwtService {
	return &jwtService{}
}

func (s *jwtService) GenerateToken(userID uuid.UUID) (string, error) {
	claim := jwt.MapClaims{}
	claim["user_id"] = userID
	
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claim)

	signedToken, err := token.SignedString(SECRET_KEY)
	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}