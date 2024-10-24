package api

import (
	"encoding/json"
	"log"
	"net/http"

	"cloud.google.com/go/firestore"
	"github.com/dledbetter123/xeddy/internal/db"
	"github.com/dledbetter123/xeddy/internal/models"
	"github.com/dledbetter123/xeddy/internal/square"
	"github.com/gorilla/mux"
)

type App struct {
	FirestoreClient *firestore.Client
}

func NewApp(firestoreClient *firestore.Client) *App {
	return &App{FirestoreClient: firestoreClient}
}

func (app *App) AddRestaurantHandler(w http.ResponseWriter, r *http.Request) {
	restaurant := models.ExampleRestaurant()

	_, _, err := app.FirestoreClient.Collection("restaurants").Add(r.Context(), restaurant)
	if err != nil {
		http.Error(w, "Failed to add restaurant", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Restaurant added successfully"))
}

func (app *App) GetRestaurantsHandler(w http.ResponseWriter, r *http.Request) {
	var restaurants []models.Restaurant

	iter := app.FirestoreClient.Collection("restaurants").Documents(r.Context())
	for {
		doc, err := iter.Next()
		if err != nil {
			break
		}
		var restaurant models.Restaurant
		doc.DataTo(&restaurant)
		restaurants = append(restaurants, restaurant)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(restaurants)
}

func OAuthCallbackHandler(w http.ResponseWriter, r *http.Request) {

	code := r.URL.Query().Get("code")
	if code == "" {
		http.Error(w, "Authorization code not provided", http.StatusBadRequest)
		return
	}

	tokenResp, err := square.ExchangeCodeForToken(code)
	if err != nil {
		http.Error(w, "Failed to exchange token: "+err.Error(), http.StatusInternalServerError)
		return
	}

	log.Printf("Access Token: %s", tokenResp.AccessToken)
	log.Printf("Merchant ID: %s", tokenResp.MerchantID)

	firestoreClient, err := db.InitializeFirestore()
	if err != nil {
		http.Error(w, "Failed to initialize Firestore: "+err.Error(), http.StatusInternalServerError)
		return
	}

	ctx := r.Context()
	err = db.StoreSquareToken(ctx, firestoreClient, tokenResp.MerchantID, tokenResp.AccessToken)
	if err != nil {
		http.Error(w, "Failed to store token in Firestore: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("OAuth successful! Token has been stored."))
}

func SetupRoutes(app *App) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/add-restaurant", app.AddRestaurantHandler).Methods("POST")
	router.HandleFunc("/restaurants", app.GetRestaurantsHandler).Methods("GET")
	router.HandleFunc("/oauth/callback", OAuthCallbackHandler)
	return router
}
