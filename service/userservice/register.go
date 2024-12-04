package userservice

import (
	"context"
	"fmt"
	"hotel_with_test/entity"
	"hotel_with_test/params"
)

func (s UserService) Register(request params.RegisterRequest) (params.RegisterResponse, error) {

	user := entity.MongoUser{
		Name:     request.Name,
		Email:    request.Email,
		Password: getMD5Hash(request.Password),
		IsAdmin:  false,
	}

	createdUser, err := s.repo.Insert(context.Background(), user)
	if err != nil {
		return params.RegisterResponse{}, fmt.Errorf("unexpected error: %w", err)
	}

	return params.RegisterResponse{User: params.UserInfo{
		Name:  createdUser.Name,
		Email: createdUser.Email,
	}}, nil
}
