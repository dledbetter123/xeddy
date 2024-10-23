package db

import (
	"context"

	"cloud.google.com/go/firestore"
)

func StoreSquareToken(ctx context.Context, client *firestore.Client, merchantID string, token string) error {
	_, err := client.Collection("merchants").Doc(merchantID).Set(ctx, map[string]interface{}{
		"square_token": token,
	})
	return err
}
