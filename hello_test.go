package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func (t *testing.T) {
		got := Hello("Oliver")
		want := "Hello, Oliver"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
