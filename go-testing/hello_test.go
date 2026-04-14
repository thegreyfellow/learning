package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	got := Hello("earth")
	want := "Hello, earth"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
