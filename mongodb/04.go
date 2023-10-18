package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
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

	filter := bson.D{{"name", "0001"}}
	replacement := util.Restaurant{Name: "Monsieur Vo", Cuisine: "Asian Fusion", Address: "Shanghai"}

	result, err := coll.ReplaceOne(context.TODO(), filter, replacement)
	if err != nil {
		panic(err)
	}
	fmt.Printf("result:%+v", result)
}
