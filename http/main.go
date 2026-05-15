package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/VitorAngelozi/quickNotes/internal/server"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load(".env", "http/.env")

	host := getenv("HOST", "127.0.0.1")
	port := getenv("PORT", "8080")
	addr := fmt.Sprintf("%s:%s", host, port)

	mux := server.NewMux()

	log.Printf("server listening on http://%s", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}

func getenv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
