package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

var Client *mongo.Client

func InitMongoDB(uri string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	serverApi := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverApi)

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		log.Printf("Error connecting to MongoDB: %v", err)
		return err
	}

	if err := client.Database("Testnet").RunCommand(ctx, bson.D{{"ping", 1}}).Err(); err != nil {
		log.Printf("Error pinging MongoDB: %v", err)
		return err
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
	Client = client

	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Printf("Failed to list databases: %v", err)
		return err
	}

	databasesExists := false
	for _, dbName := range databases {
		if dbName == "Testnet" {
			log.Println("Database 'Testnet' already exists")
			databasesExists = true
			break
		}
	}

	if !databasesExists {
		err := client.Database("Testnet").CreateCollection(ctx, "testnets")
		if err != nil {
			log.Printf("Failed to create collection: %v", err)
			return err
		}
		fmt.Println("Collection 'testnets' created successfully")

		err = client.Database("Testnet").Collection("testnets").Drop(ctx)
		if err != nil {
			log.Println("Failed to drop collection: ", err)
			return err
		}
		fmt.Println("Collection 'testnets' dropped successfully")
	} else {
		fmt.Println("Database 'Testnet' already exists")
	}

	adminCollection := client.Database("Testnet").Collection("admin")
	count, err := adminCollection.CountDocuments(ctx, bson.M{})
	if err != nil {
		log.Println("Failed to count documents")
		return err
	}

	if count == 0 {
		err := godotenv.Load()
		if err != nil {
			log.Fatalf("Error loading .env file: %v", err)
		}

		username := os.Getenv("INITADMINUSER")
		password := os.Getenv("INITADMINPASSWORD")

		hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			log.Println("Failed to hash password")
			return err
		}

		adminDoc := bson.D{
			{Key: "username", Value: string(username)},
			{Key: "password", Value: string(hashedPass)},
			{Key: "isAdmin", Value: true},
		}
		_, err = adminCollection.InsertOne(ctx, adminDoc)
		if err != nil {
			log.Println("Failed to insert admin document")
			return err
		}
		fmt.Println("Admin document inserted successfully")
	} else {
		fmt.Println("Admin document already exists")
	}

	return nil
}
