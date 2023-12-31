package main

import "testing"

func TestRepeat(t *testing.T) {
	got := Repeat(".")
	want := "....."
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
