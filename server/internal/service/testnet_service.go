package service

import (
	"context"
	"errors"
	"fmt"
	"tg-bot-server/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type TestService struct {
	db *mongo.Database
}

func NewTestService(db *mongo.Database) *TestService {
	return &TestService{db: db}
}

func (t *TestService) CreateTestnet(ctx context.Context, name, description, link, airdropdata string) error {
	testnets := &models.Testnet{}
	err := t.db.Collection("testnet").FindOne(ctx, bson.M{"nameTestnet": name, "descriptionTestnet": description, "linkTestnet": link, "dataAirdropTestnet": airdropdata}).Decode(testnets)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return errors.New("testnet not found")
		}
		return err
	}

	_, err = t.db.Collection("testnet").InsertOne(ctx, models.Testnet{Name: name, Description: description, Link: link, DataAirdrop: airdropdata})
	return err
}

func (t *TestService) CreateTestnets(ctx context.Context, names, descriptions, links, airdropdates []string) ([]string, error) {
	var newTestnets []string

	for i, name := range names {
		existingTestnet := &models.Testnet{}
		err := t.db.Collection("testnets").FindOne(ctx, bson.M{"name": name, "description": descriptions[i], "link": links[i], "dataAirdrop": airdropdates[i]}).Decode(existingTestnet)

		if err == mongo.ErrNoDocuments {
			testnet := models.Testnet{
				Name:        name,
				Description: descriptions[i],
				Link:        links[i],
				DataAirdrop: airdropdates[i],
			}

			_, err := t.db.Collection("testnets").InsertOne(ctx, testnet)
			if err != nil {
				return nil, err
			}
			newTestnets = append(newTestnets, name)
		} else if err != nil {
			return nil, err
		}
	}

	return newTestnets, nil
}

func (t *TestService) GetTestnets(ctx context.Context) ([]models.Testnet, error) {
	cursor, err := t.db.Collection("testnets").Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var tests []models.Testnet
	if err := cursor.All(ctx, &tests); err != nil {
		return nil, err
	}
	return tests, nil
}

func (t *TestService) UpdateTestnet(ctx context.Context, testnetName, newName, newDescription, newLink, newDataAirdrop string) error {
	update := bson.M{
		"$set": bson.M{
			"nameTestnet":        newName,
			"descriptionTestnet": newDescription,
			"linkTestnet":        newLink,
			"dataAirdropTestnet": newDataAirdrop,
		},
	}

	result, err := t.db.Collection("testnets").UpdateOne(ctx, bson.M{"nameTestnet": testnetName}, update)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("no testnet found with name: %s", testnetName)
	}

	return nil
}

func (t *TestService) DeleteTestnet(ctx context.Context, name string) error {
	result, err := t.db.Collection("testnet").DeleteOne(ctx, bson.M{"nameTestnet": name})
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return errors.New("testnet not found")
	}

	return nil
}
