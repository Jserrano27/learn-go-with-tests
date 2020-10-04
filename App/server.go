package main

import (
	"fmt"
	"net/http"
	"strings"
)

type PlayerStore interface {
	getPlayerScore(player string) int
	recordWin(player string)
}

type PlayerServer struct {
	store PlayerStore
}

func (p *PlayerServer) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	player := strings.TrimPrefix(req.URL.Path, "/players/")

	switch req.Method {
	case http.MethodPost:
		p.processWin(res, player)
	case http.MethodGet:
		p.showScore(res, player)
	}
}

func (p *PlayerServer) showScore(res http.ResponseWriter, player string) {
	score := p.store.getPlayerScore(player)

	if score == 0 {
		res.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(res, p.store.getPlayerScore(player))
}

func (p *PlayerServer) processWin(res http.ResponseWriter, player string) {
	res.WriteHeader(http.StatusAccepted)
	p.store.recordWin(player)
}
