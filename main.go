package main

import (
	"context"
	"fmt"
	"hotel_with_test/entity"
	"hotel_with_test/repository/mongo"
	"hotel_with_test/repository/mongo/userrepo"
	"log"
	"time"
)

func main() {

	cnf := mongo.Config{
		Host:   "localhost",
		Port:   27017,
		DBName: "hotel_db",
	}

	mongoDB, err := mongo.New(cnf)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer mongoDB.Disconnect()

	userRepo := userrepo.New(mongoDB)

	newUser := entity.User{
		Name:              "John Doe",
		Email:             "johndoe@example.com",
		EncryptedPassword: "12345678",
		IsAdmin:           false,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	insertedUser, err := userRepo.Insert(ctx, newUser)
	if err != nil {
		log.Fatalf("Failed to insert user: %v", err)
	}

	fmt.Printf("User inserted: %+v\n", insertedUser)

}
