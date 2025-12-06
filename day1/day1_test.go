package aoc25

import "testing"

func TestP1(t *testing.T) {
	t.Run("given test input", func(t *testing.T) {
		testInput := []string{
			"L68",
			"L30",
			"R48",
			"L5",
			"R60",
			"L55",
			"L1",
			"L99",
			"R14",
			"L82",
		}
		got := SolveP1(testInput)

		want := 3

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("test to submit", func(t *testing.T) {
		got := SolveP1(day1Input)

		want := 1081

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}

func TestPart2(t *testing.T) {
	t.Run("test mod", func(t *testing.T) {
		got := mod((50 - 68), 100)
		want := 82

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	helper := func(t testing.TB, got, want int) {
		t.Helper()
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	}

	rotTests := []struct {
		s    []string
		want int
	}{
		{[]string{"L50", "R50"}, 1},
		{[]string{"R50", "R50"}, 1},
		{[]string{"L50", "L50"}, 1},
		{[]string{"R50", "L50"}, 1},
		{[]string{"L150", "L50"}, 2},
		{[]string{"L150", "R50"}, 2},
		{[]string{"R150", "L50"}, 2},
		{[]string{"R150", "R50"}, 2},
		{[]string{"L50", "L1", "R1"}, 2},
		{[]string{"L48", "R198"}, 2},
		{[]string{"L50", "L300"}, 4},
		{[]string{"L50", "R101"}, 2},
		{[]string{"R49", "L98"}, 0},
		{[]string{"R49", "R1"}, 1},
		{[]string{"R49", "R1", "R1"}, 1},
		{[]string{"L50", "L400"}, 5},
		{[]string{"R50", "R400"}, 5},
		{testInput, 6},
		{day1Input, 6689},
	}

	for _, tt := range rotTests {
		t.Run("table test", func(t *testing.T) {
			t.Logf("got: %q want %d\n", tt.s, tt.want)
			helper(t, SolveP2(tt.s), tt.want)
		})
	}
}

var testInput = []string{
	"L68",
	"L30",
	"R48",
	"L5",
	"R60",
	"L55",
	"L1",
	"L99",
	"R14",
	"L82",
}
