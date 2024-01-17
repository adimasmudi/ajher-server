package database

import (
	"context"
	"log"
	"os"

	firestore "cloud.google.com/go/firestore"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
)

func FirestoreConnection() *firestore.Client {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Replace the path to your service account JSON key file
	credentialsPath := os.Getenv("FIREBASE_SERVICE_ACCOUNT_PATH")

	// Create a new context
	ctx := context.Background()

	// Use the context and service account key file to create a Firestore client
	client, err := firestore.NewClient(ctx, os.Getenv("FIRESTORE_PROJECT_ID"), option.WithCredentialsFile(credentialsPath))
	if err != nil {
		log.Fatalf("Failed to create Firestore client: %v", err)
	}

	return client
}
