package params

import "go.mongodb.org/mongo-driver/bson/primitive"

type HotelCreateRequest struct {
	Name     string   `json:"name"`
	Location string   `json:"location"`
	Rooms    []string `json:"rooms"`
	Rating   float64  `json:"rating"`
}

type HotelCreateResponse struct {
	ID       primitive.ObjectID `json:"id"`
	Name     string             `json:"name"`
	Location string             `json:"location"`
	Rating   float64            `json:"rating"`
}
