package main

import "testing"

func assertHello(t *testing.T, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestHello(t *testing.T) {
	t.Run("Test Base Hello", func(t *testing.T) {
		got := Hello("Hello dojo", "EN")
		want := "Hello dojo"
		assertHello(t, got, want)
	})

	t.Run("Test Empty Hello JP", func(t *testing.T) {
		got := Hello("", "JP")
		want := "私の道場へようこそ"
		assertHello(t, got, want)
	})

	t.Run("Test Empty Hello", func(t *testing.T) {
		got := Hello("", "EN")
		want := "Hello dojo"
		assertHello(t, got, want)
	})
}
