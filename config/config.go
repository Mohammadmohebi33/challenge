package config

import (
	"hotel_with_test/repository/mongo"
	"hotel_with_test/service/authservice"
)

type HTTPServer struct {
	Port int
}

type Config struct {
	Mongo      mongo.Config
	HTTPServer HTTPServer
	Auth       authservice.Config
}
