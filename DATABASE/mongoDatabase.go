package database

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func init() {
	err := godotenv.Load(".env") // To get env files
	// Classic error checking
	if err != nil {
		log.Fatal("There is an error in Loading the env file", err)
	}
}

// Function to get a mongo client
func GetMongoCLient() *mongo.Client {
	// Creating a client..
	client, err := mongo.Connect(options.Client().ApplyURI(os.Getenv("MONGO_STRING")))
	// classic error check
	if err != nil {
		log.Fatal("There is an error in fetching the client")
	}
	// Client Success
	log.Println("Successfully got mongoClient")
	return client
}

// Function to get a mongoCollection..
func CreateMongoCollection(client *mongo.Client, mongoDataBaseName string) *mongo.Collection {
	mongoCollenction := client.Database(mongoDataBaseName).Collection("user")

	// Created A Mongo Collection
	return mongoCollenction
}
