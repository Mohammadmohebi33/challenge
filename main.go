package main

import (
	"hotel_with_test/config"
	"hotel_with_test/delivery/httpserver"
	"hotel_with_test/delivery/httpserver/bookinghandler"
	"hotel_with_test/delivery/httpserver/hotelhandler"
	"hotel_with_test/delivery/httpserver/roomhandler"
	"hotel_with_test/delivery/httpserver/userhandler"
	"hotel_with_test/repository/mongo"
	"hotel_with_test/repository/mongo/bookingrepo"
	"hotel_with_test/repository/mongo/hotelrepo"
	"hotel_with_test/repository/mongo/roomrepo"
	"hotel_with_test/repository/mongo/userrepo"
	"hotel_with_test/service/authservice"
	"hotel_with_test/service/bookingservice"
	"hotel_with_test/service/hotelservice"
	"hotel_with_test/service/roomservice"
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
	hotelRepo := hotelrepo.New(mongoDb)
	roomRepo := roomrepo.New(mongoDb)
	bookingRepo := bookingrepo.New(mongoDb)

	authService := authservice.New(cfg.Auth)
	userService := userservice.New(authService, userRepo)
	hotelService := hotelservice.NewHotelService(hotelRepo)
	roomService := roomservice.NewHotelService(roomRepo)
	bookingService := bookingservice.NewBookingService(bookingRepo)

	userHandler := userhandler.NewUserHandler(userService, authService, userRepo)
	hotelHandler := hotelhandler.New(hotelService, authService, userRepo)
	roomHandler := roomhandler.New(roomService, authService, userRepo)
	bookingHandler := bookinghandler.New(bookingService, authService, userRepo)

	server := httpserver.NewServer(cfg, userHandler, hotelHandler, roomHandler, bookingHandler)
	server.StartServer()
}
