package config

import "hotel_with_test/repository/mongo"

type Config struct {
	mongo mongo.Config
}
