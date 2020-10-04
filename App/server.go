package main

import (
	"fmt"
	"net/http"
	"strings"
)

type PlayerStore interface {
	getPlayerScore(player string) int
}

type InMemoryPlayerStore struct{}

func (i *InMemoryPlayerStore) getPlayerScore(name string) int {
	return 20
}

type PlayerServer struct {
	store PlayerStore
}

func (p *PlayerServer) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	player := strings.TrimPrefix(req.URL.Path, "/players/")

	fmt.Fprint(res, p.store.getPlayerScore(player))
}
