package main

import (
	"encoding/json"
	"io"
)

type FileSystemPlayerStore struct {
	database io.ReadWriteSeeker
	league   League
}

func NewFileSystemPlayerStore(db io.ReadWriteSeeker) *FileSystemPlayerStore {
	db.Seek(0, 0)
	league, _ := NewLeague(db)

	return &FileSystemPlayerStore{
		database: db,
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

	f.database.Seek(0, 0)
	json.NewEncoder(f.database).Encode(f.league)
}

func (f *FileSystemPlayerStore) getLeague() League {
	f.database.Seek(0, 0)
	league, _ := NewLeague(f.database)
	return league
}
