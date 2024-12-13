package hotelservice

import (
	"context"
	"hotel_with_test/entity"
)

type HotelRepository interface {
	Insert(context.Context, entity.Hotel) (entity.Hotel, error)
	GetHotelByID(context.Context, string) (entity.Hotel, error)
	GetAllHotels(context.Context) ([]entity.Hotel, error)
	GetRoomsByHotelID(context.Context, string) ([]entity.Room, error)
}

type HotelService struct {
	repo HotelRepository
}

func NewHotelService(repo HotelRepository) HotelService {
	return HotelService{
		repo: repo,
	}
}
