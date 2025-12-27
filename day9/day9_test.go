package day9

import "testing"

func TestLargestArea(t *testing.T) {
	t.Run("example.txt", func(t *testing.T) {
		got := LargestArea("example.txt")
		want := 50

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})

	t.Run("input.txt", func(t *testing.T) {
		got := LargestArea("input.txt")
		want := 4786902990

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}

