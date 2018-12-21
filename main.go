package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/mike-mcdonald/go-trimet-gtfs/api"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	appID := os.Getenv("TRIMETAPPID")

	tick := time.Tick(30 * time.Second)
	var since int64

	client, _ := mongo.NewClient("mongodb://localhost:27017")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	client.Connect(ctx)
	cancel()
	collection := client.Database("trimet").Collection("vehicle_locations")

	for {
		select {
		case <-tick:
			v := api.RetrieveVehicles(appID, since)
			since = v.QueryTime
			for _, vehicle := range v.Vehicle {
				ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
				val, _ := bson.Marshal(vehicle)
				collection.InsertOne(ctx, val)
				cancel()
			}
		}
	}
}
