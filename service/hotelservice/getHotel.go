package hotelservice

import (
	"context"
	"hotel_with_test/entity"
)

func (s HotelService) GetHotelById(ctx context.Context, id string) (entity.Hotel, error) {
	hotel, err := s.repo.GetHotelByID(ctx, id)
	if err != nil {
		return entity.Hotel{}, err
	}
	return hotel, nil
}

func (s HotelService) GetAllHotel(ctx context.Context) ([]entity.Hotel, error) {
	hotels, err := s.repo.GetAllHotels(ctx)
	if err != nil {
		return []entity.Hotel{}, err
	}
	return hotels, nil
}

func (s HotelService) GetRooms(ctx context.Context, id string) ([]entity.Room, error) {
	rooms, err := s.repo.GetRoomsByHotelID(ctx, id)
	if err != nil {
		return []entity.Room{}, err
	}
	return rooms, nil
}
