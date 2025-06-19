package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

func GetSubTestnets(ctx context.Context, testnetName string) ([]string, error) {
	collection := client.Database("testnet").Collection("testnets")
	filter := bson.M{"name": testnetName}
	var result struct {
		SubTestnet []struct {
			SubTestnetName string `bson:"name"`
		} `bson:"subTestnet"`
	}
	err := collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		log.Printf("Failed to get sub testnets: %v", err)
		return nil, err
	}

	var subTestnets []string
	for _, subTestnet := range result.SubTestnet {
		subTestnets = append(subTestnets, subTestnet.SubTestnetName)
	}
	return subTestnets, nil
}
