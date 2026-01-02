package day4

import "testing"


func TestSolveP1(t *testing.T) {

	t.Run("example test data", func(t *testing.T) {
		got := SolveP1(example_data)

		want :=13

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})

	t.Run("example test data", func(t *testing.T) {
		got := SolveP1(day4_data)

		want :=1464

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}

func TestSolveP2(t *testing.T) {
	t.Run("example test data", func(t *testing.T) {
		got := SolveP2(example_data)

		want := 43

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("example test data", func(t *testing.T) {
		got := SolveP2(day4_data)

		want :=8409

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}
