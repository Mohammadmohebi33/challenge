package userrepo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"hotel_with_test/entity"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//	Insert(context.Context, entity.User) (entity.User, error)

func (d *DB) Insert(ctx context.Context, user entity.User) (entity.User, error) {

	collection := d.conn.Conn().Collection("users")

	_, err := collection.InsertOne(ctx, bson.M{
		"_id":      primitive.NewObjectID(),
		"name":     user.Name,
		"email":    user.Email,
		"password": user.Password,
		"isAdmin":  false,
	})
	if err != nil {
		return entity.User{}, err
	}

	// Return the inserted user
	return user, nil
}

func (d *DB) GetUerByEmail(ctx context.Context, email string) (entity.User, error) {
	var user entity.User

	collection := d.conn.Conn().Collection("users")

	filter := bson.M{"email": email}

	err := collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return entity.User{}, fmt.Errorf("user with email %s not found", email)
		}
		return entity.User{}, fmt.Errorf("failed to find user by email: %w", err)
	}

	return user, nil
}
