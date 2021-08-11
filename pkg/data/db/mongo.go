package db

import (
	"context"
	"fmt"
	"os"
	"time"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/Abhishek-More/customapi-backend/pkg/utils"
)

func Connect() *mongo.Client {
	mongo_uri := os.Getenv("URI")
	log.Fatal(mongo_uri)
	// Connect to MongoDB
	client, err := mongo.NewClient(options.Client().ApplyURI(mongo_uri))
	utils.CheckFatalError(err);
	ctx, _ := context.WithTimeout(context.Background(), 3 * time.Second)
	err = client.Connect(ctx)
	utils.CheckFatalError(err);

	fmt.Println("Connected to MongoDB!")

	return client
}

func Disconnect(client *mongo.Client) {
	ctx, _ := context.WithTimeout(context.Background(), 3 * time.Second)
	client.Disconnect(ctx)
}