package config

import "hotel_with_test/repository/mongo"

type HTTPServer struct {
	Port int
}

type Config struct {
	Mongo      mongo.Config
	HTTPServer HTTPServer
}
