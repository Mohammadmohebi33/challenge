package bookingrepo

import "hotel_with_test/repository/mongo"

type DB struct {
	conn *mongo.MongoDB
}

func New(conn *mongo.MongoDB) *DB {
	return &DB{
		conn: conn,
	}
}
