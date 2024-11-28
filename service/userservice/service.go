package userservice

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"hotel_with_test/entity"
)

type UserRepository interface {
	Insert(context.Context, entity.User) (entity.User, error)
}

type UserService struct {
	repo UserRepository
}

func New(repo UserRepository) UserService {
	return UserService{repo: repo}
}

func getMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
