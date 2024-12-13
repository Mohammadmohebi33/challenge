package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Booking struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	UserID    primitive.ObjectID `bson:"userID,omitempty" json:"user_id,omitempty"`
	RoomID    primitive.ObjectID `bson:"roomID,omitempty" json:"room_id,omitempty"`
	NumPerson int                `bson:"num_person,omitempty" json:"num_person,omitempty"`
	FromDate  time.Time          `bson:"from_date,omitempty" json:"from_date,omitempty"`
	TillDate  time.Time          `bson:"till_date,omitempty" json:"till_date,omitempty"`
	Canceled  bool               `bson:"canceled" json:"canceled"`
}
