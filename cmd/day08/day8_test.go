package day8

import (
	"testing"
)

func TestJunctionBoxes(t *testing.T) {
	t.Run("example.txt", func(t *testing.T) {
		got := JunctionBoxes("example.txt", 10)
		want := 40

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})

	t.Run("input.txt", func(t *testing.T) {
		got := JunctionBoxes("input.txt", 1000)
		want := 47040

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}

func TestJunctionBoxesP2(t *testing.T) {
	t.Run("example.txt", func(t *testing.T) {
		got := JunctionBoxesP2("example.txt")
		want := 25272

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})

	t.Run("input.txt", func(t *testing.T) {
		got := JunctionBoxesP2("input.txt")
		want := 4884971896

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}
