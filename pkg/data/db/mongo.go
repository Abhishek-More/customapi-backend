package db

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/Abhishek-More/customapi-backend/pkg/utils"
)

func Connect() *mongo.Client {
	mongo_uri := os.Getenv("URI")

	// Connect to MongoDB
	client, err := mongo.NewClient(options.Client().ApplyURI(mongo_uri))
	utils.CheckError(err);
	ctx, _ := context.WithTimeout(context.Background(), 3 * time.Second)
	err = client.Connect(ctx)
	utils.CheckError(err);

	fmt.Println("Connected to MongoDB!")

	return client
}

func Disconnect(client *mongo.Client) {
	ctx, _ := context.WithTimeout(context.Background(), 3 * time.Second)
	client.Disconnect(ctx)
}