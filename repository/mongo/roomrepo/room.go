package roomrepo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"hotel_with_test/entity"
)

func (d *DB) Insert(ctx context.Context, room entity.Room) (entity.Room, error) {

	collection := d.conn.Conn().Collection("rooms")

	_, err := collection.InsertOne(ctx, bson.M{
		"_id":       room.ID,
		"size":      room.Size,
		"basePrice": room.BasePrice,
		"price":     room.Price,
		"hotelID":   room.HotelID,
	})
	if err != nil {
		return entity.Room{}, err
	}

	return room, nil
}
