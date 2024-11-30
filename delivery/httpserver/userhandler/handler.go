package userhandler

import "hotel_with_test/service/userservice"

type UserHandler struct {
	UserService userservice.UserService
}

func NewUserHandler(userSvc userservice.UserService) UserHandler {
	return UserHandler{
		UserService: userSvc,
	}
}
