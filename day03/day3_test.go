package day3

import "testing"

func TestP1(t *testing.T) {
	t.Run("example test data", func(t *testing.T) {
		got := SolveP1(example_data)

		want := 357

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("day3 test data", func(t *testing.T) {
		got := SolveP1(day3_test_data)

		want := 17155

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}

func TestP2attempt2(t *testing.T) {
	
	t.Run("example test data", func(t *testing.T) {
		got := SolveP2attempt2(example_data)

		want := 3121910778619
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("test data", func(t *testing.T) {
		got := SolveP2attempt2(day3_test_data)

		want := 169685670469164
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}
