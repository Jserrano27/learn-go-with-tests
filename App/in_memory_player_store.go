package main

import (
	"sync"
)

type InMemoryPlayerStore struct {
	mu    sync.Mutex
	store map[string]int
}

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{
		sync.Mutex{},
		map[string]int{},
	}
}

func (i *InMemoryPlayerStore) getPlayerScore(player string) int {
	i.mu.Lock()
	defer i.mu.Unlock()

	return i.store[player]
}

func (i *InMemoryPlayerStore) recordWin(player string) {
	i.mu.Lock()
	defer i.mu.Unlock()

	i.store[player]++
}

func (i *InMemoryPlayerStore) getLeague() []Player {
	i.mu.Lock()
	defer i.mu.Unlock()

	var league []Player

	for name, wins := range i.store {
		league = append(league, Player{name, wins})
	}

	return league
}
