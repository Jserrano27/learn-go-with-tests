package main

import "testing"

func TestFileSystemStore(t *testing.T) {

	t.Run("/league from a reader", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 32}]`)

		defer cleanDatabase()

		store := NewFileSystemPlayerStore(database)

		got := store.getLeague()
		want := []Player{
			{"Cleo", 10},
			{"Chris", 32},
		}

		// read again
		got = store.getLeague()

		assertLeague(t, got, want)
	})

	t.Run("/get player score", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 32}]`)

		defer cleanDatabase()

		store := NewFileSystemPlayerStore(database)

		got := store.getPlayerScore("Cleo")
		want := 10
		assertScoreEquals(t, got, want)
	})

	t.Run("store wins for existing players", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 32}]`)

		defer cleanDatabase()

		store := NewFileSystemPlayerStore(database)
		store.recordWin("Cleo")

		got := store.getPlayerScore("Cleo")
		want := 11
		assertScoreEquals(t, got, want)
	})

	t.Run("store wins for new players", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 32}]`)

		defer cleanDatabase()

		store := NewFileSystemPlayerStore(database)
		store.recordWin("Pepper")

		got := store.getPlayerScore("Pepper")
		want := 1
		assertScoreEquals(t, got, want)
	})
}

func assertScoreEquals(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
