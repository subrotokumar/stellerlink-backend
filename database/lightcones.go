package database

import (
	"context"
	"log"
	"time"

	"github.com/subrotokumar/stellerlink-backend/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (db *DB) GetLightCone(id int) *model.LightCone {
	lightConesCoollection := db.client.Database(dbName).Collection("lightcones")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	log.Printf("Search for id %v", id)
	filter := bson.M{"id": id}
	var lightCone model.LightCone
	err := lightConesCoollection.FindOne(ctx, filter).Decode(&lightCone)

	if err == mongo.ErrNoDocuments {
		log.Printf("No document found for id %v", id)
		return nil
	}

	if err != nil {
		log.Fatalf("Error retrieving LightCone: %v", err)
	}

	log.Printf("lightCone found: %v", lightCone.ID)
	return &lightCone
}

func (db *DB) GetLightCones() []*model.LightCone {
	lightConeCollection := db.client.Database(dbName).Collection("lightcones")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var lightCones []*model.LightCone
	cursor, err := lightConeCollection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
		return nil
	}
	if err = cursor.All(context.TODO(), &lightCones); err != nil {
		log.Fatal(err)
		return nil
	}

	return lightCones
}
