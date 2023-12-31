package auth

import (
	"errors"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type Service interface {
	GenerateToken(UserID uuid.UUID) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
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

func (s *jwtService) ValidateToken(encodedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error){
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if ! ok {
			return nil, errors.New("Invalid Token")
		}
		return []byte(SECRET_KEY), nil
	})

	if err != nil {
		return token, err
	}
	return token, nil
}