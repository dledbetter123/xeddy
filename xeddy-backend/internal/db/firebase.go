package db

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

func InitializeFirestore() (*firestore.Client, error) {
	ctx := context.Background()
	sa := option.WithCredentialsFile("internal/db/xeddy_firebase_credentials.json")

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
