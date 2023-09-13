package database

import (
	"context"
	"log"
	"time"

	"github.com/subrotokumar/stellerlink-backend/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (db *DB) GetRelic(id int) *model.Relic {
	characterCollection := db.client.Database(dbName).Collection("relics")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	log.Printf("Search for id %v", id)
	filter := bson.M{"_id": id}
	var relic model.Relic
	err := characterCollection.FindOne(ctx, filter).Decode(&relic)

	if err == mongo.ErrNoDocuments {
		// Handle the case where no document was found (e.g., return an error or handle it as needed).
		log.Printf("No document found for id %v", id)
		return nil // or return an error, depending on your use case
	}

	if err != nil {
		log.Fatalf("Error retrieving relic: %v", err)
	}

	log.Printf("Relic found: %v", relic.ID)
	return &relic
}

func (db *DB) GetRelics() []*model.Relic {
	relicCollection := db.client.Database(dbName).Collection("relics")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var relics []*model.Relic
	cursor, err := relicCollection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
		return nil
	}
	if err = cursor.All(context.TODO(), &relics); err != nil {
		log.Fatal(err)
		return nil
	}

	return relics
}
