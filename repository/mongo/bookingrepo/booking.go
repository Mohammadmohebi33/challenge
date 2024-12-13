package bookingrepo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"hotel_with_test/entity"
)

func (d *DB) InsertBooking(ctx context.Context, booking entity.Booking) (entity.Booking, error) {

	collection := d.conn.Conn().Collection("booking")

	_, err := collection.InsertOne(ctx, bson.M{
		"useID":      booking.UserID,
		"roomID":     booking.RoomID,
		"num_person": booking.NumPerson,
		"from_date":  booking.FromDate,
		"till_date":  booking.TillDate,
		"canceled":   booking.Canceled,
	})
	if err != nil {
		return entity.Booking{}, err
	}

	return booking, nil
}
