package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"pratice/mongodb/util"
)

func main() {
	client := util.Conn()

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	coll := client.Database("test").Collection("restaurants")

	id, _ := primitive.ObjectIDFromHex("652ce59b5e9f5238f5dfeba3")
	filter := bson.D{{"_id", id}}
	update := bson.D{{"$set", bson.D{{"avg_rating", 4.4}}}}

	result, err := coll.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		panic(err)
	}
	fmt.Printf("result:%+v", result)
}
