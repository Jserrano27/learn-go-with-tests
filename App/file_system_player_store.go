package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type FileSystemPlayerStore struct {
	database *json.Encoder
	league   League
}

func NewFileSystemPlayerStore(file *os.File) (*FileSystemPlayerStore, error) {
	file.Seek(0, 0)

	info, err := file.Stat()
	if err != nil {
		return nil, fmt.Errorf("problem reading file info from file %s, %v", file.Name(), err)
	}

	if info.Size() == 0 {
		file.Write([]byte(`[]`))
		file.Seek(0, 0)
	}

	league, err := NewLeague(file)

	if err != nil {
		return nil, fmt.Errorf("problem loading player store from file %s, %v", file.Name(), err)
	}

	return &FileSystemPlayerStore{
		database: json.NewEncoder(file),
		league:   league,
	}, nil
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
