package main

import (
	"hotel_with_test/config"
	"hotel_with_test/delivery/httpserver"
	"hotel_with_test/delivery/httpserver/userhandler"
	"hotel_with_test/repository/mongo"
	"hotel_with_test/repository/mongo/userrepo"
	"hotel_with_test/service/authservice"
	"hotel_with_test/service/userservice"
)

func main() {

	mongoConfig := mongo.Config{
		Host:   "localhost",
		Port:   27017,
		DBName: "hotel_db",
	}

	authConfig := authservice.Config{
		SignKey:               config.JwtSignKey,
		AccessExpirationTime:  config.AccessTokenExpireDuration,
		RefreshExpirationTime: config.RefreshTokenExpireDuration,
		AccessSubject:         config.AccessTokenSubject,
		RefreshSubject:        config.RefreshTokenSubject,
	}

	cfg := config.Config{
		Mongo: mongoConfig,
		HTTPServer: config.HTTPServer{
			Port: 8080,
		},
		Auth: authConfig,
	}

	mongoDb, err := mongo.New(cfg.Mongo)
	if err != nil {
		panic(err)
	}

	userRepo := userrepo.New(mongoDb)
	authService := authservice.New(cfg.Auth)
	userService := userservice.New(authService, userRepo)
	userHandler := userhandler.NewUserHandler(userService)

	server := httpserver.NewServer(cfg, userHandler)
	server.StartServer()

}
