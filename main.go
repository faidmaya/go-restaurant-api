package main

import (
	"log"
	"os"

	"restaurant-api/database"
	"restaurant-api/routers"
)

func main() {
	// Warning jika belum set DATABASE_URL (lokal)
	if os.Getenv("DATABASE_URL") == "" {
		log.Println("Warning: DATABASE_URL not set. Set it before running in production.")
	}

	// Koneksi database
	db, err := database.Connect()
	if err != nil {
		log.Fatal("database connect:", err)
	}
	defer db.Close()

	// Setup router
	r := routers.SetupRouter(db)

	// Ambil PORT dari Railway
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // default untuk lokal
	}

	// Start server
	if err := r.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
