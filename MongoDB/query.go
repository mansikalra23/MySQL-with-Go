package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Trainee struct {
	ID    primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name  string             `json:"name,omitempty" bson:"name,omitempty`
	Batch string             `json:"batch,omitempty" bson:"batch,omitempty`
}

var client *mongo.Client

func main() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully Connected!")
	defer client.Disconnect(ctx)

	collection := client.Database("try").Collection("Trainees")

	filter := bson.D{primitive.E{Key: "name", Value: "Mansi"}}

	var trainee Trainee

	err = collection.FindOne(context.TODO(), filter).Decode(&trainee)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Trainee is : ", trainee.Name, trainee.Batch)

}
