package hotelhandler

import (
	"hotel_with_test/repository/mongo/userrepo"
	"hotel_with_test/service/authservice"
	"hotel_with_test/service/hotelservice"
)

type HotelHandler struct {
	hotelService hotelservice.HotelService
	authService  authservice.Service
	UserStore    *userrepo.DB
}

func New(hotelSvc hotelservice.HotelService, authSvc authservice.Service, userStore *userrepo.DB) HotelHandler {
	return HotelHandler{
		hotelService: hotelSvc,
		authService:  authSvc,
		UserStore:    userStore,
	}
}
