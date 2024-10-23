package square

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type MerchantProfile struct {
	ID           string `json:"id"`
	BusinessName string `json:"business_name"`
	Country      string `json:"country"`
	Language     string `json:"language_code"`
}

func GetMerchantProfile(accessToken string) (*MerchantProfile, error) {
	url := "https://connect.squareup.com/v2/merchants/me"

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var profile MerchantProfile
	err = json.NewDecoder(resp.Body).Decode(&profile)
	if err != nil {
		return nil, err
	}

	return &profile, nil
}

type OrderRequest struct {
	LocationID string     `json:"location_id"`
	LineItems  []LineItem `json:"line_items"`
}

type LineItem struct {
	Name           string `json:"name"`
	Quantity       string `json:"quantity"`
	BasePriceMoney Money  `json:"base_price_money"`
}

type Money struct {
	Amount   int64  `json:"amount"`
	Currency string `json:"currency"`
}

func PlaceOrder(accessToken, locationID string, orderRequest *OrderRequest) error {
	url := fmt.Sprintf("https://connect.squareup.com/v2/orders")

	orderRequest.LocationID = locationID
	jsonBody, _ := json.Marshal(orderRequest)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to place order: %s", resp.Status)
	}

	return nil
}
