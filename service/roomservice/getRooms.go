package roomservice

import (
	"context"
	"hotel_with_test/entity"
)

func (s RoomService) GetAllRooms(ctx context.Context) ([]entity.Room, error) {
	rooms, err := s.repo.GetAll(ctx)
	if err != nil {
		return []entity.Room{}, err
	}
	return rooms, nil
}
