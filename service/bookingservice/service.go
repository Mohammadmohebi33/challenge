package bookingservice

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"hotel_with_test/entity"
	"hotel_with_test/params"
)

type BookingRepository interface {
	InsertBooking(ctx context.Context, booking entity.Booking) (entity.Booking, error)
}

type BookingService struct {
	repo BookingRepository
}

func NewBookingService(repo BookingRepository) BookingService {
	return BookingService{
		repo: repo,
	}
}

func (s BookingService) Book(ctx context.Context, userID primitive.ObjectID, roomID primitive.ObjectID, req params.BookRoomParams) (entity.Booking, error) {

	booking := entity.Booking{
		UserID:    userID,
		RoomID:    roomID,
		NumPerson: req.NumPerson,
		FromDate:  req.FromDate,
		TillDate:  req.TillDate,
		Canceled:  false,
	}

	resp, err := s.repo.InsertBooking(ctx, booking)
	if err != nil {
		return entity.Booking{}, err
	}

	return resp, nil
}
