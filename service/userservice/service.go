package userservice

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"hotel_with_test/entity"
)

type UserRepository interface {
	Insert(context.Context, entity.User) (entity.User, error)
	GetUerByEmail(context.Context, string) (entity.User, error)
}

type AuthGenerator interface {
	CreateAccessToken(user entity.User) (string, error)
	CreateRefreshToken(user entity.User) (string, error)
}

type UserService struct {
	auth AuthGenerator
	repo UserRepository
}

func New(auth AuthGenerator, repo UserRepository) UserService {
	return UserService{
		auth: auth,
		repo: repo,
	}
}

func getMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
