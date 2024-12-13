package roomservice

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"hotel_with_test/entity"
	"hotel_with_test/params"
)

func (s RoomService) CreateRoom(ctx context.Context, req params.CreateRoomRequest) (params.CreateRoomResponse, error) {

	hotelID, err := primitive.ObjectIDFromHex(req.HotelID)
	room := entity.Room{
		ID:        primitive.NewObjectID(),
		Size:      req.Size,
		BasePrice: req.BasePrice,
		Price:     req.Price,
		HotelID:   hotelID,
	}
	resp, err := s.repo.Insert(ctx, room)
	if err != nil {
		return params.CreateRoomResponse{}, err
	}

	err = s.repo.UpdateHotel(ctx, resp.HotelID.Hex(), resp.ID.Hex())
	if err != nil {
		return params.CreateRoomResponse{}, err
	}

	return params.CreateRoomResponse{
		ID:        resp.ID,
		Size:      resp.Size,
		BasePrice: resp.BasePrice,
		Price:     resp.Price,
		HotelID:   resp.HotelID,
	}, nil
}
