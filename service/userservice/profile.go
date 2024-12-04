package userservice

import (
	"context"
	"hotel_with_test/params"
)

func (s UserService) Profile(ctx context.Context, req params.ProfileRequest) (params.ProfileResponse, error) {

	user, err := s.repo.GetUserByID(ctx, req.UserID)
	if err != nil {
		return params.ProfileResponse{}, err
	}
	return params.ProfileResponse{Name: user.Name}, nil
}
