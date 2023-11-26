package configs

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const uri = "mongodb://myuser:mypassword@localhost:27017"

func disconnectDb(err error, client *mongo.Client, context context.Context) {
	if err = client.Disconnect(context); err != nil {
		panic(err)
	}
}

func StartDb() {
	serverApi := options.ServerAPI((options.ServerAPIVersion1))
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverApi)

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		panic(err)
	}

	defer disconnectDb(err, client, ctx)

	var result bson.M
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{Key: "ping", Value: 1}}).Decode(&result); err != nil {
		panic(err)
	}

	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
}
