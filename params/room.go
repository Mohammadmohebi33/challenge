package params

import "go.mongodb.org/mongo-driver/bson/primitive"

type CreateRoomRequest struct {
	Size      string  `json:"size"`
	BasePrice float64 ` json:"basePrice"`
	Price     float64 `json:"price"`
	HotelID   string  `json:"hotelID"`
}

type CreateRoomResponse struct {
	ID        primitive.ObjectID `json:"id,omitempty"`
	Size      string             `json:"size"`
	BasePrice float64            ` json:"basePrice"`
	Price     float64            `json:"price"`
	HotelID   primitive.ObjectID `json:"hotelID"`
}
