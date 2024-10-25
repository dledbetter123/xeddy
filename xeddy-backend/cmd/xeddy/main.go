package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/dledbetter123/xeddy/internal/api"
	"github.com/dledbetter123/xeddy/internal/db"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load() // Looks for a .env file in the project root
	if err != nil {
		log.Printf("Error loading .env file: %v", err)
	}
	firestoreClient, err := db.InitializeFirestore()
	if err != nil {
		log.Fatalf("Failed to initialize Firestore: %v", err)
	}

	app := api.NewApp(firestoreClient)
	router := api.SetupRoutes(app)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Printf("Server starting on port %s\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), router))
}
