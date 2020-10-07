package main

import (
	"encoding/json"
	"io"
	"os"
)

type FileSystemPlayerStore struct {
	database io.Writer
	league   League
}

func NewFileSystemPlayerStore(db *os.File) *FileSystemPlayerStore {
	db.Seek(0, 0)
	league, _ := NewLeague(db)

	return &FileSystemPlayerStore{
		database: &tape{db},
		league:   league,
	}
}

func (f *FileSystemPlayerStore) getPlayerScore(name string) int {
	player := f.league.Find(name)

	if player != nil {
		return player.Wins
	}

	return 0
}

func (f *FileSystemPlayerStore) recordWin(name string) {
	player := f.league.Find(name)

	if player != nil {
		player.Wins++
	} else {
		f.league = append(f.league, Player{name, 1})
	}

	json.NewEncoder(f.database).Encode(f.league)
}

func (f *FileSystemPlayerStore) getLeague() League {
	return f.league
}
