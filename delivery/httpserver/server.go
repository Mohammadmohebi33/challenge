package httpserver

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"hotel_with_test/config"
	"hotel_with_test/delivery/httpserver/hotelhandler"
	"hotel_with_test/delivery/httpserver/roomhandler"
	"hotel_with_test/delivery/httpserver/userhandler"
)

type Server struct {
	config       config.Config
	userHandler  userhandler.UserHandler
	hotelHandler hotelhandler.HotelHandler
	roomHandler  roomhandler.RoomHandler
	Router       *fiber.App
}

func NewServer(config config.Config, userHandler userhandler.UserHandler, hotelHandler hotelhandler.HotelHandler, roomHandler roomhandler.RoomHandler) Server {

	app := fiber.New(fiber.Config{})

	return Server{
		config:       config,
		userHandler:  userHandler,
		hotelHandler: hotelHandler,
		roomHandler:  roomHandler,
		Router:       app,
	}
}

func (receiver Server) StartServer() {
	// Routes
	receiver.Router.Get("/health-check", receiver.healthCheck)
	receiver.userHandler.SetRoutes(receiver.Router)
	receiver.hotelHandler.SetRoute(receiver.Router)
	receiver.roomHandler.SetRoute(receiver.Router)

	// Start server
	address := fmt.Sprintf(":%d", receiver.config.HTTPServer.Port)
	fmt.Printf("start echo server on %s\n", address)
	if err := receiver.Router.Listen(address); err != nil {
		fmt.Println("router start error", err)
	}
}
