package square

//  xeddy/square/square.go

import (
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

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

	scope := "MERCHANT_PROFILE_READ PAYMENTS_WRITE" // basic need

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
