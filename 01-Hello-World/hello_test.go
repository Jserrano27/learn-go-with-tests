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
		got := sayHello("Chris")
		want := "Hello, Chris!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying 'Hello, World!' when name is not informed", func(t *testing.T) {
		got := sayHello("")
		want := "Hello, World!"
		assertCorrectMessage(t, got, want)
	})
}
