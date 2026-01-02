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

func TestLargestRectangle(t *testing.T) {
	t.Run("example.txt", func(t *testing.T) {
		got := MakeLargestRectangle("example.txt")
		want := 24

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})

	t.Run("input.txt", func(t *testing.T) {
		got := MakeLargestRectangle("input.txt")
		want := 1571016172

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}
