package main

import "testing"

func TestHello(t *testing.T) {
	got := sayHello("Peter")
	want := "Hello, Peter!"

	if got != want {
		t.Errorf("got: %q, but wanted: %q", got, want)
	}
}
