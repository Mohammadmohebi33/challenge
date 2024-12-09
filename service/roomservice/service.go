package roomservice

import (
	"context"
	"hotel_with_test/entity"
)

type RoomRepository interface {
	Insert(context.Context, entity.Room) (entity.Room, error)
}

type RoomService struct {
	repo RoomRepository
}

func NewHotelService(repo RoomRepository) RoomService {
	return RoomService{
		repo: repo,
	}
}
