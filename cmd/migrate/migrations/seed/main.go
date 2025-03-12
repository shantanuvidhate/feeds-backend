package main

import (
	"log"

	"github.com/joho/godotenv"

	"github.com/shantanuvidhate/feeds-backend/internal/db"
	"github.com/shantanuvidhate/feeds-backend/internal/env"
	"github.com/shantanuvidhate/feeds-backend/internal/store"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	addr := env.GetString("DB_ADDR", "postgres://user:password@localhost:5432/feeds?sslmode=disable")
	conn, err := db.New(addr, 5, 5, "15m")
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	store := store.NewStorage(conn)

	db.Seed(&store, conn)
}
