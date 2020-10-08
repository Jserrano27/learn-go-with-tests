package poker

import "testing"

func TestFileSystemStore(t *testing.T) {
	t.Run("/get player score", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 32}]`)

		defer cleanDatabase()

		store, err := NewFileSystemPlayerStore(database)
		assertNoError(t, err)

		got := store.getPlayerScore("Cleo")
		want := 10
		assertScoreEquals(t, got, want)
	})

	t.Run("store wins for existing players", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 32}]`)

		defer cleanDatabase()

		store, err := NewFileSystemPlayerStore(database)
		assertNoError(t, err)

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

		store, err := NewFileSystemPlayerStore(database)
		assertNoError(t, err)

		store.recordWin("Pepper")

		got := store.getPlayerScore("Pepper")
		want := 1
		assertScoreEquals(t, got, want)
	})

	t.Run("works with an empty file", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, ``)
		defer cleanDatabase()

		_, err := NewFileSystemPlayerStore(database)
		assertNoError(t, err)
	})

	t.Run("sort league by score", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 32}]`)

		defer cleanDatabase()

		store, err := NewFileSystemPlayerStore(database)
		assertNoError(t, err)

		got := store.getLeague()
		want := []Player{
			{"Chris", 32},
			{"Cleo", 10},
		}

		assertLeague(t, got, want)

		// read again
		got = store.getLeague()
		assertLeague(t, got, want)
	})
}

func assertScoreEquals(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func assertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("didn't expect an error but got one: %v", err)
	}
}
