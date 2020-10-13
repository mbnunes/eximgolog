package eximgolog

import (
	context "context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// InsertLogLine - insere o struct LogLine no mongodb
func InsertLogLine(logline LogLine) {

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))

	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)

	collection := client.Database("eximgolog").Collection("logs")
	insertResult, err := collection.InsertOne(ctx, logline)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted post with ID:", insertResult.InsertedID)

	defer client.Disconnect(ctx)
}

// FindLogLine - Procura o e-mail deacordo com o seu mailid
func FindLogLine(mailid string) {

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))

	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	//	err = client.Connect(ctx)

	collection := client.Database("eximgolog").Collection("logs")
	err = collection.FindOne(ctx, bson.D{"mailid": 25})

	if err != nil {
		log.Fatal(err)
	}

	defer client.Disconnect(ctx)
}
