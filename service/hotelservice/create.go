package hotelservice

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"hotel_with_test/entity"
	"hotel_with_test/params"
)

func (s HotelService) Create(ctx context.Context, req params.HotelCreateRequest) (params.HotelCreateResponse, error) {

	hotel := entity.Hotel{
		ID:       primitive.NewObjectID(),
		Name:     req.Name,
		Location: req.Location,
		Rating:   req.Rating,
	}
	hotelInserted, err := s.repo.Insert(ctx, hotel)
	if err != nil {
		return params.HotelCreateResponse{}, err
	}

	return params.HotelCreateResponse{
		ID:       hotelInserted.ID,
		Name:     hotelInserted.Name,
		Location: hotelInserted.Location,
		Rating:   hotelInserted.Rating,
	}, err
}
