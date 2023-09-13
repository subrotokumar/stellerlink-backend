package database

import (
	"context"
	"log"
	"strings"
	"time"

	"github.com/subrotokumar/stellerlink-backend/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (db *DB) GetCharacter(id int) *model.Character {
	characterCollection := db.client.Database(dbName).Collection("characters")
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
	characterCollection := db.client.Database(dbName).Collection("characters")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var characters []*model.Character
	cursor, err := characterCollection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
		return nil
	}
	if err = cursor.All(context.TODO(), &characters); err != nil {
		log.Fatal(err)
		return nil
	}

	return characters
}

func (db *DB) AddCharacter(input *model.CharacterInput) *model.Character {
	characterCollection := db.client.Database(dbName).Collection("characters")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if input.Images == nil {
		name := strings.ReplaceAll(strings.ToLower(input.Name), " ", "_")
		input.Images = &model.ImageInput{
			Profile:     "/images/characters/" + name + "/profile.webp",
			Splash:      "/images/characters/" + name + "/splash.webp",
			Transparent: "/images/characters/" + name + "/transparent.webp",
		}
	}

	insert, err := characterCollection.InsertOne(ctx, &input)

	if err != nil {
		log.Fatal(err)
		return nil
	}

	insert.InsertedID.(primitive.ObjectID).Hex()

	return db.GetCharacter(input.ID)
}
