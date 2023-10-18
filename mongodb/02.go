package main

import (
	"context"
	"fmt"
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
	newRestaurant := &util.Restaurant{Name: "0001", Cuisine: "aaaa"}
	result, err := coll.InsertOne(context.TODO(), newRestaurant)
	if err != nil {
		panic(err)
	}
	fmt.Printf("result:%+v", result)
}
