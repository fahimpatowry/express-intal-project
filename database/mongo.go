package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

// ConnectMongo connects to MongoDB
func ConnectMongo(uri string) error {
	var err error
	client, err = mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		return err
	}

	// Ping to test connection
	if err := client.Ping(ctx, nil); err != nil {
		return err
	}

	fmt.Println("Connected to MongoDB!")
	return nil
}

// GetCollection returns a MongoDB collection
func GetCollection(dbName, collectionName string) *mongo.Collection {
	return client.Database(dbName).Collection(collectionName)
}

// InsertSampleUser inserts a test user
func InsertSampleUser(dbName string) {
	collection := GetCollection(dbName, "users")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	user := bson.D{
		{"name", "Fahim"},
		{"role", "Developer"},
		{"experience", 3.4},
	}

	result, err := collection.InsertOne(ctx, user)
	if err != nil {
		log.Fatalf("Insert failed: %v", err)
	}

	fmt.Println("Inserted document ID:", result.InsertedID)
}
