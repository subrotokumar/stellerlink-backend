package database

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/subrotokumar/stellerlink-backend/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DB struct {
	client *mongo.Client
}

func Connect() *DB {
	connectString := os.Getenv("ConnectionString")

	clientOptions := options.Client().ApplyURI(connectString)

	// Create a context with a timeout.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect to the MongoDB server.
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Ping the MongoDB server to verify the connection.
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	return &DB{
		client: client,
	}
}

func (db *DB) GetCharacter(id int) *model.Character {
	characterCollection := db.client.Database("stellerlink").Collection("characters")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	log.Printf("Search for id %v", id)
	filter := bson.M{"id": id}
	var character model.Character
	err := characterCollection.FindOne(ctx, filter).Decode(&character)

	if err == mongo.ErrNoDocuments {
		// Handle the case where no document was found (e.g., return an error or handle it as needed).
		log.Printf("No document found for id %v", id)
		return nil // or return an error, depending on your use case
	}

	if err != nil {
		log.Fatalf("Error retrieving character: %v", err)
	}

	log.Printf("Character found: %v", character.ID)
	return &character

}

func (db *DB) GetCharacters() []*model.Character {
	characterCollection := db.client.Database("stellerlink").Collection("characters")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var characters []*model.Character
	cursor, err := characterCollection.Find(ctx, bson.D{})
	log.Printf(cursor.Current.String())
	if err != nil {
		log.Fatal(err)
	}
	if err = cursor.All(context.TODO(), &characters); err != nil {
		panic(err)
	}

	return characters
}
