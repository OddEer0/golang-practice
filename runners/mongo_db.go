package runners

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	Id string
	Login string
	Password string
}

func MongoDbRun() {
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(""))
	if err != nil {
		panic(err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		panic(err)
	}

	fmt.Print("Connect mongodb")

	collection := client.Database("test").Collection("users")
}
