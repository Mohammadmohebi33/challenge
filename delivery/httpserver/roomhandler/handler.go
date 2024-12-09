package roomhandler

import (
	"hotel_with_test/repository/mongo/userrepo"
	"hotel_with_test/service/authservice"
	"hotel_with_test/service/roomservice"
)

type RoomHandler struct {
	roomService roomservice.RoomService
	authService authservice.Service
	UserStore   *userrepo.DB
}

func New(roomSvc roomservice.RoomService, authSvc authservice.Service, userStore *userrepo.DB) RoomHandler {
	return RoomHandler{
		roomService: roomSvc,
		authService: authSvc,
		UserStore:   userStore,
	}
}
