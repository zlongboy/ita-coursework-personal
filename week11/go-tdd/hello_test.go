package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	got := Hello("Zack")
	want := "Hello, Zack"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
