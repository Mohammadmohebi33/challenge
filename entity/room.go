package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Room struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Size      string             `bson:"size" json:"size"`
	BasePrice float64            `bson:"basePrice" json:"basePrice"`
	Price     float64            `bson:"price" json:"price"`
	HotelID   primitive.ObjectID `bson:"hotelID,omitempty" json:"hotelId,omitempty"`
}
