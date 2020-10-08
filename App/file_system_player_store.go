package main

import (
	"encoding/json"
	"os"
)

type FileSystemPlayerStore struct {
	database *json.Encoder
	league   League
}

func NewFileSystemPlayerStore(file *os.File) *FileSystemPlayerStore {
	file.Seek(0, 0)
	league, _ := NewLeague(file)

	return &FileSystemPlayerStore{
		database: json.NewEncoder(file),
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

	f.database.Encode(f.league)
}

func (f *FileSystemPlayerStore) getLeague() League {
	return f.league
}
