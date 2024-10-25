package db

import (
	"context"
	"log"
	"os"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
)

func InitializeFirestore() (*firestore.Client, error) {
	ctx := context.Background()
	err := godotenv.Load() // Looks for a .env file in the project root
	if err != nil {
		log.Printf("Error loading .env file: %v", err)
	}
	firebaseCredentialsJSON := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")
	if firebaseCredentialsJSON == "" {
		log.Fatal("GOOGLE_APPLICATION_CREDENTIALS environment variable not set")
	}

	sa := option.WithCredentialsJSON([]byte(firebaseCredentialsJSON))

	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Printf("Failed to initialize Firebase app: %v", err)
		return nil, err
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Printf("Failed to initialize Firestore client: %v", err)
		return nil, err
	}

	return client, nil
}
