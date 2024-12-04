package userrepo

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"hotel_with_test/entity"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//	Insert(context.Context, entity.User) (entity.User, error)

func (d *DB) Insert(ctx context.Context, user entity.MongoUser) (entity.MongoUser, error) {

	collection := d.conn.Conn().Collection("users")

	_, err := collection.InsertOne(ctx, bson.M{
		"_id":      primitive.NewObjectID(),
		"name":     user.Name,
		"email":    user.Email,
		"password": user.Password,
		"isAdmin":  false,
	})
	if err != nil {
		return entity.MongoUser{}, err
	}

	// Return the inserted user
	return user, nil
}

func (d *DB) GetUerByEmail(ctx context.Context, email string) (entity.MongoUser, error) {
	var user entity.MongoUser

	collection := d.conn.Conn().Collection("users")

	filter := bson.M{"email": email}

	err := collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return entity.MongoUser{}, fmt.Errorf("user with email %s not found", email)
		}
		return entity.MongoUser{}, fmt.Errorf("failed to find user by email: %w", err)
	}

	return user, nil
}

func (d *DB) GetUserByID(ctx context.Context, userID interface{}) (entity.MongoUser, error) {
	id, ok := userID.(string)
	if !ok {
		return entity.MongoUser{}, errors.New("invalid userID type for MongoDB")
	}

	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return entity.MongoUser{}, err
	}
	collection := d.conn.Conn().Collection("users")
	filter := bson.M{"_id": oid}
	var user entity.MongoUser
	err = collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return entity.MongoUser{}, fmt.Errorf("user with id %s not found", id)
		}
		return entity.MongoUser{}, fmt.Errorf("failed to find user by id: %w", err)
	}
	return user, nil

}
