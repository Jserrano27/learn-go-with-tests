package poker

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// PlayerStore stores score information about players
type PlayerStore interface {
	getPlayerScore(player string) int
	recordWin(player string)
	getLeague() League
}

// Player stores a name with a number of wins
type Player struct {
	Name string
	Wins int
}

// PlayerServer is a HTTP interface for player information
type PlayerServer struct {
	store PlayerStore
	http.Handler
}

// NewPlayerServer creates a PlayerServer with routing configured
func NewPlayerServer(store PlayerStore) *PlayerServer {
	p := new(PlayerServer)

	p.store = store
	router := http.NewServeMux()

	router.Handle("/players/", http.HandlerFunc(p.playerHandler))
	router.Handle("/league", http.HandlerFunc(p.leagueHandler))

	p.Handler = router

	return p
}

func (p *PlayerServer) playerHandler(res http.ResponseWriter, req *http.Request) {
	player := strings.TrimPrefix(req.URL.Path, "/players/")

	switch req.Method {
	case http.MethodPost:
		p.processWin(res, player)
	case http.MethodGet:
		p.showScore(res, player)
	}
}

func (p *PlayerServer) leagueHandler(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "application/json")
	json.NewEncoder(res).Encode(p.store.getLeague())
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
