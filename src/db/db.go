package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB mongo.Database

func ConnectToDb() {
    client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))

    if err != nil {
        log.Fatal(err)
    }

	ctx, _ := context.WithTimeout(context.Background(), 24*time.Hour)
    err = client.Connect(ctx)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Connected to MongoDB...")
	DB = *client.Database("goFiber")
}
