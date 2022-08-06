package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Zack", "")
		want := "Hello, Zack"
		errorMsg(t, got, want)
	})
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Zack", "Spanish")
		want := "Hola, Zack"
		errorMsg(t, got, want)
	})
	t.Run("in French", func(t *testing.T) {
		got := Hello("Zack", "French")
		want := "Bonjour, Zack"
		errorMsg(t, got, want)
	})
	t.Run("say 'Hello, world' when empty string is passed in", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world"
		errorMsg(t, got, want)
	})
}

// Testing helpers
func errorMsg(t testing.TB, got, want string) { // testing.TB is an interface for testing.T and testing.B
	t.Helper() // Indicates this function is a helper and directs failure log to line of actual failure (function call)
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
