package database

import (
	"log"

	"github.com/dgraph-io/ristretto"
)

var Cache *ristretto.Cache

func InitCache() {
	var err error
	Cache, err = ristretto.NewCache(&ristretto.Config{
		NumCounters: 1e7, // Number of keys to track frequency of (10M).
		MaxCost:     110, // Max number of lobbies i want to hold concurrently
		BufferItems: 64,  // Number of keys per Get buffer.
	})

	if err != nil {
		log.Fatalf("Failed to initialize Ristretto cache: %v", err)
	}
}
