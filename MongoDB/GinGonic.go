package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
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

	r := gin.Default()

	// Routes
	r.GET("/trainees", getTrainees)

	r.Run(":8000")
}

func getTrainees(c *gin.Context) {
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

	c.JSON(http.StatusOK, gin.H{"data": trainees})
}
