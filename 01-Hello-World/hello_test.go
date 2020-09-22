package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q, but wanted %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := SayHello("Chris", "English")
		want := "Hello, Chris!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying 'Hello, World!' when name is not informed", func(t *testing.T) {
		got := SayHello("", "English")
		want := "Hello, World!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := SayHello("Roberto", "Spanish")
		want := "Hola, Roberto!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := SayHello("Lucien", "French")
		want := "Bonjour, Lucien!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Italian", func(t *testing.T) {
		got := SayHello("Bambino", "Italian")
		want := "Ciao, Bambino!"
		assertCorrectMessage(t, got, want)
	})
}
