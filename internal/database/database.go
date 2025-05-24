package database

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/supabase-community/supabase-go"
)

var Supabase *supabase.Client

func Init() {
	// Load environment variables
	_ = godotenv.Load()

	url := os.Getenv("SUPABASE_URL")
	key := os.Getenv("SUPABASE_KEY")

	if url == "" || key == "" {
		log.Fatal("SUPABASE_URL and SUPABASE_KEY must be set")
	}

	// Initialize Supabase client
	client, err := supabase.NewClient(url, key, nil)
	if err != nil {
		log.Fatalf("Failed to initialize Supabase client: %v", err)
	}

	Supabase = client
}
