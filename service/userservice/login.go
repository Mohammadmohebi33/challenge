package userservice

import (
	"context"
	"fmt"
	"hotel_with_test/params"
)

func (s UserService) Login(req params.LoginRequest) (params.LoginResponse, error) {

	user, err := s.repo.GetUerByEmail(context.Background(), req.Email)
	if err != nil {
		return params.LoginResponse{}, err
	}

	fmt.Println(user.Name, user.Email, user.Password)
	fmt.Println(getMD5Hash(req.Password))

	if user.Password != getMD5Hash(req.Password) {
		return params.LoginResponse{}, fmt.Errorf("username or password isn't correct")
	}

	accessToken, err := s.auth.CreateAccessToken(user)
	if err != nil {
		return params.LoginResponse{}, fmt.Errorf("unexpected error: %w", err)
	}

	refreshToken, err := s.auth.CreateRefreshToken(user)
	if err != nil {
		return params.LoginResponse{}, fmt.Errorf("unexpected error: %w", err)
	}

	return params.LoginResponse{
		User: params.UserInfo{
			ID:    user.ID,
			Email: user.Email,
			Name:  user.Name,
		},
		Tokens: params.Tokens{
			AccessToken:  accessToken,
			RefreshToken: refreshToken},
	}, nil

}
