package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Service interface {
	Health() map[string]string
}

type service struct {
	db *mongo.Client
}

var (
	// host        = os.Getenv("DB_HOST")
	// port        = os.Getenv("DB_PORT")
	DB_USER     = os.Getenv("DB_USERNAME")
	DB_PASSWORD = os.Getenv("DB_ROOT_PASSWORD")
	DB_NAME     = os.Getenv("DB_NAME")
	//database = os.Getenv("DB_DATABASE")
)

func New() Service {

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(fmt.Sprintf("mongodb+srv://%s:%s@cluster0.ifstf4g.mongodb.net/%s?retryWrites=true&w=majority", DB_USER, DB_NAME, DB_PASSWORD)))

	if err != nil {
		log.Fatal(err)

	}
	return &service{
		db: client,
	}
}

func (s *service) Health() map[string]string {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	err := s.db.Ping(ctx, nil)
	if err != nil {
		log.Fatalf(fmt.Sprintf("db down: %v", err))
	}

	return map[string]string{
		"message": "It's healthy",
	}
}
