package userservice

import (
	"context"
	"fmt"
	"hotel_with_test/entity"
	"hotel_with_test/params"
)

func (receiver UserService) Register(request params.RegisterRequest) (params.RegisterResponse, error) {

	user := entity.User{
		ID:                0,
		Name:              request.Name,
		Email:             request.Email,
		EncryptedPassword: getMD5Hash(request.Password),
		IsAdmin:           false,
	}

	createdUser, err := receiver.repo.Insert(context.Background(), user)
	if err != nil {
		return params.RegisterResponse{}, fmt.Errorf("unexpected error: %w", err)
	}

	return params.RegisterResponse{User: params.UserInfo{
		ID:    createdUser.ID,
		Name:  createdUser.Name,
		Email: createdUser.Email,
	}}, nil
}
