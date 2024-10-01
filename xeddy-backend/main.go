package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/dledbetter123/xeddy/models"
	"github.com/gorilla/mux"
	"google.golang.org/api/option"
)

var client *firestore.Client

func ExampleRestaurant() models.Restaurant {
	return models.Restaurant{
		ID:               1,
		Name:             "Gusto Italian Grill",
		Category:         "Italian",
		ImageURL:         "https://example.com/image.jpg",
		LogoURL:          "https://example.com/logo.jpg",
		Location:         "1234 Culinary Blvd, Taste City, TC 56789",
		DescriptionLong:  "Gusto Italian Grill offers a rich taste of Italy with a diverse menu featuring traditional dishes, handpicked wines, and a cozy ambiance perfect for any occasion.",
		DescriptionShort: "Traditional Italian cuisine in the heart of Taste City.",
		Email:            "contact@gustoitalian.com",
		Phone:            "555-1234",
		OddDates:         []time.Time{time.Now(), time.Now().AddDate(0, 0, 1)}, // Keep valid timestamps
		Hours: map[string]models.TimeObject{ // Use strings for open/close times
			"Monday":  {Open: "10:00:00", Close: "22:00:00"},
			"Tuesday": {Open: "10:00:00", Close: "22:00:00"},
		},
		Closed: false,
		Menu: []models.Item{
			{ID: 1, Name: "Spaghetti Carbonara", Description: "Creamy pasta with pancetta, egg, and Parmesan", Price: 15.99},
			{ID: 2, Name: "Margherita Pizza", Description: "Classic pizza with fresh tomatoes, mozzarella, and basil", Price: 12.50},
		},
		Rewards: []int64{100, 200}, // Example points
		Deals: []models.Deal{
			{ID: 1, Description: "20% off on Wednesdays", ValidUntil: time.Now().AddDate(0, 1, 0)}, // Keep valid timestamps
		},
		OrderSystems: []string{"Toast", "Square"},
		Reviews: []models.Review{
			{ID: 1, Rating: 5, Text: "Outstanding service and food quality!"},
		},
	}
}

type App struct {
	FirestoreClient *firestore.Client
}

func (app *App) addRestaurantHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	var restaurant models.Restaurant = ExampleRestaurant()
	// log.Println("Restaurant JSON :", restaurant)
	// err := json.NewDecoder(r.Body).Decode(&restaurant)
	// if err != nil {
	// 	http.Error(w, "Invalid request payload", http.StatusBadRequest)
	// 	return
	// }

	// showing the json representation of the restaurant for debugging
	jsonData, err := json.Marshal(restaurant)
	if err == nil {
		log.Println("Restaurant JSON After Marshaling:", string(jsonData))
	} else {
		log.Println("Error marshalling restaurant:", err)
	}

	_, _, err = app.FirestoreClient.Collection("restaurants").Add(ctx, restaurant)
	if err != nil {
		http.Error(w, "Failed to add restaurant", http.StatusInternalServerError)
		log.Printf("Failed to add restaurant: %v", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Restaurant added successfully"))
}

func (app *App) getRestaurantsHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	var restaurants []models.Restaurant

	iter := app.FirestoreClient.Collection("restaurants").Documents(ctx)
	for {
		doc, err := iter.Next()
		if err != nil {
			break
		}

		var restaurant models.Restaurant

		doc.DataTo(&restaurant)
		restaurants = append(restaurants, restaurant)
	}

	// Marshal the list of restaurants to JSON and write to the response
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(restaurants)
	if err != nil {
		http.Error(w, "Failed to encode restaurants", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func initializeFirestore() *firestore.Client {
	ctx := context.Background()
	sa := option.WithCredentialsFile("xeddy-native-app.json")

	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalf("Failed to initialize Firebase app: %v", err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalf("Failed to initialize Firestore client: %v", err)
	}

	return client
}

func (app *App) testFirestoreConnectionHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	_, _, err := app.FirestoreClient.Collection("test").Add(ctx, map[string]interface{}{
		"testField": "testValue",
	})
	if err != nil {
		http.Error(w, "Failed to write test document", http.StatusInternalServerError)
		log.Printf("Failed to write test document: %v", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Test document added successfully"))
}

func main() {
	firestoreClient := initializeFirestore()

	app := &App{FirestoreClient: firestoreClient}

	router := mux.NewRouter()

	router.HandleFunc("/add-restaurant", app.addRestaurantHandler).Methods("POST")
	router.HandleFunc("/test-post", app.testFirestoreConnectionHandler).Methods("POST")
	router.HandleFunc("/restaurants", app.getRestaurantsHandler).Methods("GET")

	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

// func getRestaurants(w http.ResponseWriter, r *http.Request) {
// 	ctx := context.Background()
// 	iter := client.Collection("restaurants").Documents(ctx)
// 	for {
// 		doc, err := iter.Next()
// 		if err != nil {
// 			break
// 		}
// 		log.Printf("Restaurant: %v", doc.Data())
// 	}
// }

// func addRestaurant(ctx context.Context, client *firestore.Client, restaurant models.Restaurant) {
// 	_, _, err := client.Collection("restaurants").Add(ctx, restaurant)
// 	if err != nil {
// 		log.Fatalf("Failed to add restaurant: %v", err) // Add detailed error logging here
// 	} else {
// 		log.Println("Restaurant added successfully!")
// 	}
// }
