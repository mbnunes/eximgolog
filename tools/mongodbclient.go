package tools

import (
	context "context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	con  *mongo.Client
	err  error
	ctx  context.Context
	coll *mongo.Collection
}

// ConnectMongoDb - é uma função que gera a conexão com o Mongo sempre que chamada.
func (c *MongoDB) ConnectMongoDb() {
	c.con, c.err = mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))

	if c.err != nil {
		log.Fatal(c.err)
	}

	c.ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	c.err = c.con.Connect(c.ctx)
	c.coll = c.con.Database("eximgolog").Collection("logs")
}

func (c *MongoDB) CloseConnection() {
	c.con.Disconnect(c.ctx)
}

// InsertLogLine - insere o struct LogLine no mongodb
func (c *MongoDB) InsertLogLine(logline LogLine) {

	insertResult, err := c.coll.InsertOne(c.ctx, logline)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted post with ID:", insertResult.InsertedID)

}

// FindLogLine - Procura o e-mail deacordo com o seu mailid
func (c *MongoDB) FindLogLine(dados FindForm) {

	cur, err := c.coll.Find(c.ctx, bson.M{"mailid": dados.Mailid})

	if err != nil {
		log.Fatal(err)
	}

	defer cur.Close(c.ctx)

	for cur.Next(c.ctx) {
		// To decode into a struct, use cursor.Decode()
		var result LogLine
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%+v\n", result)
	}

}
