package roomrepo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func (d *DB) UpdateHotel(ctx context.Context, hotelID string, roomID string) error {

	hotelId, err := primitive.ObjectIDFromHex(hotelID)
	if err != nil {
		return err
	}

	roomId, err := primitive.ObjectIDFromHex(roomID)
	if err != nil {
		return err
	}
	collection := d.conn.Conn().Collection("hotels")

	update := bson.M{
		"$push": bson.M{
			"rooms": roomId,
		},
	}

	_, err = collection.UpdateOne(ctx, bson.M{"_id": hotelId}, update)
	if err != nil {
		return err
	}
	return nil
}

func (d *DB) GetAll(ctx context.Context) ([]entity.Room, error) {

	var rooms []entity.Room
	collection := d.conn.Conn().Collection("rooms")

	findOptions := options.Find()

	cursor, err := collection.Find(ctx, bson.M{}, findOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var room entity.Room
		if err := cursor.Decode(&room); err != nil {
			return nil, err
		}
		rooms = append(rooms, room)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return rooms, nil
}
