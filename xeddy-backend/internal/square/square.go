package square

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"

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

func ExchangeCodeForToken(code string) (*TokenResponse, error) {
	err := godotenv.Load() // Looks for a .env file in the project root
	if err != nil {
		log.Printf("Error loading .env file: %v", err)
	}
	clientID := os.Getenv("SQUARE_APPLICATION_ID")
	clientSecret := os.Getenv("SQUARE_CLIENT_SECRET")
	var redirectURI string = ""
	if os.Getenv("DEPLOY_TYPE") == "SANDBOX" || os.Getenv("DEPLOY_TYPE") == "DEV" {
		redirectURI = os.Getenv("SQUARE_SANDBOX_REDIRECT_URL")
	} else {
		redirectURI = os.Getenv("SQUARE_REDIRECT_URL")
	}
	url := "https://connect.squareup.com/oauth2/token"

	body := map[string]string{
		"client_id":     clientID,
		"client_secret": clientSecret,
		"code":          code,
		"grant_type":    "authorization_code",
		"redirect_uri":  redirectURI,
	}

	jsonBody, _ := json.Marshal(body)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var tokenResp TokenResponse
	err = json.NewDecoder(resp.Body).Decode(&tokenResp)
	if err != nil {
		return nil, err
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
