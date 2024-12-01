package authservice

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"hotel_with_test/entity"
	"os"
	"strings"
	"time"
)

type Config struct {
	SignKey               string
	AccessExpirationTime  time.Duration
	RefreshExpirationTime time.Duration
	AccessSubject         string
	RefreshSubject        string
}

type Service struct {
	config Config
}

func New(cfg Config) Service {
	return Service{
		config: cfg,
	}
}

func (s Service) CreateAccessToken(user entity.User) (string, error) {
	return s.createToken(user.ID, s.config.AccessSubject, s.config.AccessExpirationTime)
}

func (s Service) CreateRefreshToken(user entity.User) (string, error) {
	return s.createToken(user.ID, s.config.RefreshSubject, s.config.RefreshExpirationTime)
}

func (s Service) createToken(userID int, subject string, expireDuration time.Duration) (string, error) {

	claims := jwt.MapClaims{
		"id":  userID,
		"exp": time.Now().Add(expireDuration).Unix(),
		"sub": subject,
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := accessToken.SignedString([]byte(s.config.SignKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (s Service) ParseToken(bearerToken string) (jwt.MapClaims, error) {

	tokenStr := strings.Replace(bearerToken, "Bearer ", "", 1)

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			fmt.Println("invalid signing method", token.Header["alg"])
			return nil, fmt.Errorf("invalid signing method")
		}
		secret := os.Getenv("JWT_SECRET")
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("invalid claims")
	}
	return claims, nil
}
