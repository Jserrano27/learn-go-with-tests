package main

import (
	"log"
	"net/http"
)

func main() {
	store := &InMemoryPlayerStore{}
	server := &PlayerServer{store}

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("Could not listen on port 5000: %v", err)
	}
}
