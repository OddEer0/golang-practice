package main

import (
	"fmt"
	"github.com/rs/xid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

func main() {
	fmt.Println("Hello")
	uid := xid.New().String()
	fmt.Println(uid)
	id, err := primitive.ObjectIDFromHex(uid)
	if err != nil {
		log.Fatal(err)
	}
	filter := bson.M{"id": id}
	fmt.Println(filter)
}
