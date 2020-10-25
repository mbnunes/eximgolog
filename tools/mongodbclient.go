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
func FindLogLine(dados FindForm) {

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))

	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)

	collection := client.Database("eximgolog").Collection("logs")
	cur, err := collection.Find(ctx, bson.M{"data": dados.Data, "horario": dados.Horario, "mailid": dados.Mailid, "tipo": dados.Tipo})

	if err != nil {
		log.Fatal(err)
	}

	defer cur.Close(ctx)

	for cur.Next(ctx) {
		// To decode into a struct, use cursor.Decode()
		var result LogLine
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%+v\n", result)
	}

	defer client.Disconnect(ctx)
}
