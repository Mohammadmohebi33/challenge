package hotelservice

import (
	"context"
	"hotel_with_test/entity"
)

type HotelRepository interface {
	Insert(context.Context, entity.Hotel) (entity.Hotel, error)
}

type HotelService struct {
	repo HotelRepository
}

func NewHotelService(repo HotelRepository) HotelService {
	return HotelService{
		repo: repo,
	}
}
