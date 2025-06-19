package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func ConnectMongoDB(uri string) {
	var err error
	client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatalf("Failed to connected MONGODB: %v", err)
	}
}

func GetTestnets(ctx context.Context) ([]string, error) {
	collection := client.Database("testnet").Collection("testnets")
	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Printf("Failed to get testnets: %v", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	var testnets []string
	for cursor.Next(ctx) {
		var result map[string]any
		if err := cursor.Decode(&result); err != nil {
			log.Printf("Failed to decode testnet: %v", err)
			return nil, err
		}
		if testnetName, ok := result["name"].(string); ok {
			testnets = append(testnets, testnetName)
		}
	}

	return testnets, nil
}
