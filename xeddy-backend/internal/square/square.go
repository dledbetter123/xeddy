package square

//  xeddy/square/square.go

import (
	"bytes"
	"context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	"cloud.google.com/go/firestore"
	"github.com/dledbetter123/xeddy/internal/session"
	"github.com/joho/godotenv"
)

// func InitializeAuth() {

// }

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresAt   string `json:"expires_at"`
	MerchantID  string `json:"merchant_id"`
}

type CatalogObject struct {
	ID       string `json:"id"`
	Type     string `json:"type"`
	ItemData struct {
		Name        string   `json:"name"`
		Description string   `json:"description"`
		ImageIDs    []string `json:"image_ids,omitempty"`
		Variations  []struct {
			ItemVariationData struct {
				PriceMoney struct {
					Amount int64 `json:"amount"`
				} `json:"price_money"`
			} `json:"item_variation_data"`
		} `json:"variations"`
	} `json:"item_data,omitempty"`
	ImageData struct {
		URL     string `json:"url"`
		Caption string `json:"caption,omitempty"`
	} `json:"image_data,omitempty"`
}

type MenuItem struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int64  `json:"price_money"`
	ImageURL    string `json:"image_url,omitempty"`
}

type CatalogResponse struct {
	Objects []CatalogObject `json:"objects"`
}

func GetMenuItems(ctx context.Context, firestoreClient *firestore.Client, merchantID string) ([]MenuItem, error) {
	// Fetch the merchant's square token from Firestore
	doc, err := firestoreClient.Collection("merchants").Doc(merchantID).Get(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve token from Firestore: %v", err)
	}
	token, ok := doc.Data()["square_token"].(string)
	if !ok {
		return nil, fmt.Errorf("square_token not found for merchant ID: %s", merchantID)
	}

	// Define Square Catalog API endpoint
	var baseURL string
	if os.Getenv("DEPLOY_TYPE") == "SANDBOX" || os.Getenv("DEPLOY_TYPE") == "DEV" {
		baseURL = "https://connect.squareupsandbox.com/v2/catalog/list?types=ITEM,IMAGE"
	} else {
		baseURL = "https://connect.squareup.com/v2/catalog/list"
	}

	// Prepare the HTTP request to fetch catalog items and images
	req, err := http.NewRequest("GET", baseURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	// HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to retrieve catalog: %s", resp.Status)
	}

	// Parse response into CatalogResponse structure
	var catalogResp CatalogResponse
	err = json.NewDecoder(resp.Body).Decode(&catalogResp)
	if err != nil {
		return nil, fmt.Errorf("failed to decode catalog response: %v", err)
	}

	// Map to hold image URLs by image ID
	imageMap := make(map[string]string)
	for _, object := range catalogResp.Objects {
		// log.Printf("Object Type: %s", object.Type)
		if object.Type == "IMAGE" {
			// log.Printf("Image ID: %s, Image URL: %s", object.ID, object.ImageData.URL)
			imageMap[object.ID] = object.ImageData.URL
		}
	}
	// log.Printf("Length of image map: %d", len(imageMap))
	// for id, url := range imageMap {
	// 	log.Printf("Image ID: %s, Image URL: %s", id, url)
	// }

	var menuItems []MenuItem
	for _, object := range catalogResp.Objects {
		if object.Type == "ITEM" {
			var price int64
			if len(object.ItemData.Variations) > 0 {
				price = object.ItemData.Variations[0].ItemVariationData.PriceMoney.Amount
			}

			var imageURL string
			if len(object.ItemData.ImageIDs) > 0 {
				imageURL = imageMap[object.ItemData.ImageIDs[0]]
			}

			menuItems = append(menuItems, MenuItem{
				ID:          object.ID,
				Name:        object.ItemData.Name,
				Description: object.ItemData.Description,
				Price:       price,
				ImageURL:    imageURL,
			})
		}
	}

	return menuItems, nil
}

func generateAuthState() (string, error) {
	bytes := make([]byte, 16)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func GenerateAuthURL(w http.ResponseWriter, r *http.Request) (string, error) {
	// for local
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file: %v", err)
	}

	authState, err := generateAuthState()
	if err != nil {
		return "", err
	}

	session, _ := session.Store.Get(r, "auth-session")
	session.Values["auth_state"] = authState
	session.Save(r, w)

	clientID := os.Getenv("SQUARE_APPLICATION_ID")
	if clientID == "" {
		return "", fmt.Errorf("SQUARE_APPLICATION_ID is not set")
	}

	var redirectURI string
	if os.Getenv("DEPLOY_TYPE") == "SANDBOX" || os.Getenv("DEPLOY_TYPE") == "DEV" {
		redirectURI = os.Getenv("SQUARE_SANDBOX_REDIRECT_URL")
	} else {
		redirectURI = os.Getenv("SQUARE_REDIRECT_URL")
	}

	if redirectURI == "" {
		return "", fmt.Errorf("redirect URI is not set")
	}

	scope := "MERCHANT_PROFILE_READ ITEMS_READ PAYMENTS_WRITE" // basic need

	var baseURL string
	if os.Getenv("DEPLOY_TYPE") == "SANDBOX" || os.Getenv("DEPLOY_TYPE") == "DEV" {
		baseURL = "https://connect.squareupsandbox.com/oauth2/authorize"
	} else {
		baseURL = "https://connect.squareup.com/oauth2/authorize"
	}
	params := url.Values{}
	params.Add("client_id", clientID)
	params.Add("response_type", "code")
	params.Add("scope", scope)
	params.Add("redirect_uri", redirectURI)
	params.Add("state", authState) // CSRF protection state

	log.Printf("Generated Square OAuth URL: %s?%s", baseURL, params.Encode())
	return fmt.Sprintf("%s?%s", baseURL, params.Encode()), nil
}

func ExchangeCodeForToken(code string) (*TokenResponse, error) {
	// load environment variables (for local dev, might not be needed in production)
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file: %v", err)
	}

	// Gather necessary environment variables
	clientID := os.Getenv("SQUARE_APPLICATION_ID")
	clientSecret := os.Getenv("SQUARE_CLIENT_SECRET")
	log.Printf("client ID: %s", clientID)
	log.Printf("client secret: %s", clientSecret)
	var redirectURI, url string

	// Choose sandbox or production based on deployment type
	if os.Getenv("DEPLOY_TYPE") == "SANDBOX" || os.Getenv("DEPLOY_TYPE") == "DEV" {
		redirectURI = os.Getenv("SQUARE_SANDBOX_REDIRECT_URL")
		url = "https://connect.squareupsandbox.com/oauth2/token"
	} else {
		redirectURI = os.Getenv("SQUARE_REDIRECT_URL")
		url = "https://connect.squareup.com/oauth2/token"
	}
	log.Printf("redirect URI: %s", redirectURI)

	body := map[string]string{
		"client_id":     clientID,
		"client_secret": clientSecret,
		"code":          code,
		"grant_type":    "authorization_code",
		"redirect_uri":  redirectURI,
	}

	jsonBody, _ := json.Marshal(body)
	log.Printf("Request body for token exchange: %s", string(jsonBody))

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	// HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// log  entire response body for debugging
	var responseBytes bytes.Buffer
	_, err = responseBytes.ReadFrom(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}
	log.Printf("Square token exchange response: %s", responseBytes.String())

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to exchange token: received status %s", resp.Status)
	}

	var tokenResp TokenResponse
	err = json.NewDecoder(&responseBytes).Decode(&tokenResp)
	if err != nil {
		return nil, fmt.Errorf("failed to decode token response: %v", err)
	}

	if tokenResp.AccessToken == "" || tokenResp.MerchantID == "" {
		return nil, fmt.Errorf("failed to retrieve access token or merchant ID from Square response")
	}

	return &tokenResp, nil
}

// func HandleOAuthCallback(w http.ResponseWriter, r *http.Request) {
// 	code := r.URL.Query().Get("code")
// 	token, err := exchangeCodeForToken(code)
// 	if err != nil {
// 		http.Error(w, "Failed to exchange token", http.StatusInternalServerError)
// 		return
// 	}

// 	db.SaveToken(token)
// 	// Redirect or respond to indicate success
// }
