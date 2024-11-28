package userrepo

import (
	"context"
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
		"password": user.EncryptedPassword,
		"isAdmin":  false,
	})
	if err != nil {
		return entity.User{}, err
	}

	// Return the inserted user
	return user, nil
}
