package day5

import "testing"

func TestSolveP1(t *testing.T) {
	t.Run("test example", func(t *testing.T) {

		got := SolveP1("example.txt")

		want := 3

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})

	t.Run("test input", func(t *testing.T) {

		got := SolveP1("input.txt")

		want := 840

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}
func TestSolveP2(t *testing.T) {
	t.Run("test example", func(t *testing.T) {

		got := SolveP2("example.txt")

		want := 14

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})

	t.Run("test overlap same start", func(t *testing.T) {
		var test_data = []FreshPair{
			{3, 7},
			{3, 9},
		}

		got := HandleCleaning(test_data)
		want := 7

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})

	t.Run("test overlap", func(t *testing.T) {
		var test_data = []FreshPair{
			{3, 9},
			{4, 7},
			{5, 7},
		}

		got := HandleCleaning(test_data)
		want := 7

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})

	t.Run("test weird thing ", func(t *testing.T) {
		var test_data = []FreshPair{
			{3, 9},
			{4, 7},
			{4, 8},
			{5, 6},
		}

		got := HandleCleaning(test_data)
		want := 7

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("test more complicated thing ", func(t *testing.T) {
		var test_data = []FreshPair{
			{3, 9},
			{4, 7},
			{4, 8},
			{4, 9},
			{4, 12},
			{5, 6},
		}

		got := HandleCleaning(test_data)
		want := 10

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})

	t.Run("test more complicated p2", func(t *testing.T) {
		var test_data = []FreshPair{
			{1, 4},
			{2, 5},
			{8, 10},
			{8, 9},
			{8, 8},
		}

		got := HandleCleaning(test_data)
		want := 8

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("test value of 1 same size as g1 end ", func(t *testing.T) {
		var test_data = []FreshPair{
			{1, 4},
			{4, 4},
		}

		got := HandleCleaning(test_data)
		want := 4

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("test input", func(t *testing.T) {

		got := SolveP2("input.txt")

		want := 359913027576322

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}
