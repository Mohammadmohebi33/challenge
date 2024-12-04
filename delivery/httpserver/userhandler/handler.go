package userhandler

import (
	"hotel_with_test/repository/mongo/userrepo"
	"hotel_with_test/service/authservice"
	"hotel_with_test/service/userservice"
)

type UserHandler struct {
	UserService userservice.UserService
	AuthService authservice.Service
	UserStore   *userrepo.DB
}

func NewUserHandler(
	userSvc userservice.UserService,
	authSvc authservice.Service,
	userStore *userrepo.DB,
) UserHandler {
	return UserHandler{
		UserService: userSvc,
		AuthService: authSvc,
		UserStore:   userStore,
	}
}
