package main

import (
	"log"
	"net/http"

	"github.com/dledbetter123/xeddy/internal/api"
	"github.com/dledbetter123/xeddy/internal/db"
)

func main() {
	firestoreClient, err := db.InitializeFirestore()
	if err != nil {
		log.Fatalf("Failed to initialize Firestore: %v", err)
	}

	app := api.NewApp(firestoreClient)
	router := api.SetupRoutes(app)
	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
