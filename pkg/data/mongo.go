package db

import (
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	// "go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/Abhishek-More/customapi-backend/pkg/utils"
)

func Connect() *mongo.Client {
	mongo_uri := os.Getenv("URI")

	// Connect to MongoDB
	client, err := mongo.NewClient(options.Client().ApplyURI(mongo_uri))
	utils.CheckError(err);
	ctx, _ := context.WithTimeout(context.Background(), 10 * time.Second)
	err = client.Connect(ctx)
	utils.CheckError(err);

	defer client.Disconnect(ctx)
	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	utils.CheckError(err)

	print(databases)
	print("Connected to MongoDB!")

	return client


}