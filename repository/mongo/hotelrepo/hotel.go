package hotelrepo

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func (d *DB) GetHotelByID(ctx context.Context, id string) (entity.Hotel, error) {
	var hotel entity.Hotel
	collection := d.conn.Conn().Collection("hotels")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return hotel, fmt.Errorf("invalid id format: %w", err)
	}

	filter := bson.M{"_id": objectID}
	err = collection.FindOne(ctx, filter).Decode(&hotel)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return hotel, fmt.Errorf("hotel not found: %w", err)
		}
		return hotel, fmt.Errorf("failed to fetch hotel: %w", err)
	}

	return hotel, nil
}

func (d *DB) GetAllHotels(ctx context.Context) ([]entity.Hotel, error) {
	var hotels []entity.Hotel
	collection := d.conn.Conn().Collection("hotels")

	findOptions := options.Find()

	cursor, err := collection.Find(ctx, bson.M{}, findOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var hotel entity.Hotel
		if err := cursor.Decode(&hotel); err != nil {
			return nil, err
		}
		hotels = append(hotels, hotel)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return hotels, nil
}
