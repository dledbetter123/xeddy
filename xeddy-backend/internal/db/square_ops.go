package db

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
)

func StoreSquareToken(ctx context.Context, client *firestore.Client, merchantID string, token string) error {
	if merchantID == "" {
		return fmt.Errorf("merchantID is empty; cannot store token")
	}

	_, err := client.Collection("merchants").Doc(merchantID).Set(ctx, map[string]interface{}{
		"square_token": token,
	})
	return err
}
