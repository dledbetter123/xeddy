package main

import (
	"context"
	"log"
	"net/http"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/gorilla/mux"
	"google.golang.org/api/option"
)

var client *firestore.Client

func main() {
	ctx := context.Background()
	sa := option.WithCredentialsFile("path/to/serviceAccountKey.json")

	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalf("Error initializing Firebase app: %v", err)
	}

	client, err = app.Firestore(ctx)
	if err != nil {
		log.Fatalf("Error initializing Firestore client: %v", err)
	}
	defer client.Close()

	router := mux.NewRouter()
	router.HandleFunc("/restaurants", getRestaurants).Methods("GET")
	router.HandleFunc("/restaurants", addRestaurant).Methods("POST")

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func getRestaurants(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	iter := client.Collection("restaurants").Documents(ctx)
	for {
		doc, err := iter.Next()
		if err != nil {
			break
		}
		log.Printf("Restaurant: %v", doc.Data())
	}
}

func addRestaurant(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	_, _, err := client.Collection("restaurants").Add(ctx, map[string]interface{}{
		"name":   "New Restaurant",
		"points": 200,
	})
	if err != nil {
		http.Error(w, "Failed to add restaurant", http.StatusInternalServerError)
		return
	}
	w.Write([]byte("Restaurant added successfully"))
}
