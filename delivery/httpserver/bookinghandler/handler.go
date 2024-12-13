package bookinghandler

import (
	"hotel_with_test/repository/mongo/userrepo"
	"hotel_with_test/service/authservice"
	"hotel_with_test/service/bookingservice"
)

type BookingHandler struct {
	bookingService bookingservice.BookingService
	authService    authservice.Service
	UserStore      *userrepo.DB
}

func New(bookingSvc bookingservice.BookingService, authSvc authservice.Service, userStore *userrepo.DB) BookingHandler {
	return BookingHandler{
		bookingService: bookingSvc,
		authService:    authSvc,
		UserStore:      userStore,
	}
}
