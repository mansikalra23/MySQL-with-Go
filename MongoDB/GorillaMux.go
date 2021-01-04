package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
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
var trainees []Trainee

func main() {
	c, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = c.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully Connected!")
	defer c.Disconnect(ctx)

	client = c

	router := mux.NewRouter()
	router.HandleFunc("/trainees", getTrainees).Methods("GET")
	http.ListenAndServe(":8000", router)
}

func getTrainees(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	collection := client.Database("try").Collection("Trainees")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

	result, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	defer result.Close(ctx)

	for result.Next(ctx) {
		var trainee Trainee
		err := result.Decode(&trainee)
		if err != nil {
			panic(err.Error())
		}
		trainees = append(trainees, trainee)
	}

	json.NewEncoder(w).Encode(trainees)
}
