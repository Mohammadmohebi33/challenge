package hotelrepo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"hotel_with_test/entity"
)

func (d *DB) Insert(ctx context.Context, hotel entity.Hotel) (entity.Hotel, error) {
	user := ctx.Value("user").(entity.MongoUser)
	collection := d.conn.Conn().Collection("hotels")

	_, err := collection.InsertOne(ctx, bson.M{
		"_id":      hotel.ID,
		"name":     hotel.Name,
		"location": hotel.Location,
		"rooms":    hotel.Rooms,
		"rating":   hotel.Rating,
		"owner":    user.ID,
	})
	if err != nil {
		return entity.Hotel{}, err
	}

	return hotel, nil
}
