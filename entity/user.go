package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type MongoUser struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Name     string             `bson:"name"`
	Email    string             `bson:"email"`
	Password string             `bson:"password"`
	IsAdmin  bool               `bson:"isAdmin"`
}
