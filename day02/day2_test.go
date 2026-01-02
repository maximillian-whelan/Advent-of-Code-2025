package day2

import "testing"

func TestP1(t *testing.T) {

	tabletest := []struct {
		value string
		got   bool
		want  bool
	}{
		{"123123", IsDuplicate("123123"), true},
		{"11", IsDuplicate("11"), true},
		{"22", IsDuplicate("22"), true},
		{"321", IsDuplicate("321"), false},
		{"14", IsDuplicate("14"), false},
	}

	helperBool := func(t testing.TB, got, want bool) {
		t.Helper()
		if got != want {
			t.Errorf("got %t want %t", got, want)
		}
	}

	for _, tt := range tabletest {
		t.Run("test duplicate code", func(t *testing.T) {

			if tt.got != tt.want {
				t.Logf("value %q, got %t, want %t", tt.value, tt.got, tt.want)
				helperBool(t, tt.got, tt.want)
			}

		})

	}
	
	tt_Find_dups_in_range := []struct {
		value string
		got   int
		want  int
	}{
		{"1188511880-1188511890", FindDupsInRange("1188511880","1188511890"), 1188511885},
		{"11-22", FindDupsInRange("11","22"), 33},
		{"95-115", FindDupsInRange("95","115"), 99},
		{"998-1012", FindDupsInRange("998","1012"), 1010},
		{"222220-222224", FindDupsInRange("222220","222224"), 222222},
		{"1698522-1698528", FindDupsInRange("1698522","1698528"),0},
		{"446443-446449", FindDupsInRange("446443","446449"),446446},
		{"38593856-38593862", FindDupsInRange("38593856","38593862"), 38593859},
	}


	helperInt := func(t testing.TB, got, want int) {
		t.Helper()
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	}

	for _, tt := range tt_Find_dups_in_range {
		t.Run("test find dups in range code", func(t *testing.T) {

			if tt.got != tt.want {
				t.Logf("value %q, got %d, want %d", tt.value, tt.got, tt.want)
				helperInt(t, tt.got, tt.want)
			}

		})

	}

	t.Run("test P1 with example data", func(t *testing.T) {
		got := SolveP1(example_data)
		want := 1227775554

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})

	t.Run("test P1 with my input", func(t *testing.T) {
		got := SolveP1(my_aoc_test_data)
		want := 31210613313

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}

func TestP2(t *testing.T) {
	tabletest := []struct {
		value string
		got   bool
		want  bool
	}{
		{"123123", IsDuplicateP2("123123"), true},
		{"11", IsDuplicateP2("11"), true},
		{"22", IsDuplicateP2("22"), true},
		{"321", IsDuplicateP2("321"), false},
		{"14", IsDuplicateP2("14"), false},
		{"999", IsDuplicateP2("999"), true},
		{"565656", IsDuplicateP2("565656"), true},
	}

	helperBool := func(t testing.TB, got, want bool) {
		t.Helper()
		if got != want {
			t.Errorf("got %t want %t", got, want)
		}
	}

	for _, tt := range tabletest {
		t.Run("test duplicate code", func(t *testing.T) {

			if tt.got != tt.want {
				t.Logf("value %q, got %t, want %t", tt.value, tt.got, tt.want)
				helperBool(t, tt.got, tt.want)
			}

		})
	}

	tt_Find_dups_in_range := []struct {
		value string
		got   int
		want  int
	}{
		{"1188511880-1188511890", FindDupsInRangeP2("1188511880","1188511890"), 1188511885},
		{"11-22", FindDupsInRangeP2("11","22"), 33},
		{"95-115", FindDupsInRangeP2("95","115"), 99+111},
		{"998-1012", FindDupsInRangeP2("998","1012"), 1010+999},
		{"222220-222224", FindDupsInRangeP2("222220","222224"), 222222},
		{"1698522-1698528", FindDupsInRangeP2("1698522","1698528"),0},
		{"446443-446449", FindDupsInRangeP2("446443","446449"),446446},
		{"38593856-38593862", FindDupsInRangeP2("38593856","38593862"), 38593859},
		{"824824821-824824827", FindDupsInRangeP2("824824821","824824827"), 824824824},
		{"2121212118-2121212124", FindDupsInRangeP2("2121212118","2121212124"), 2121212121},
	}


	helperInt := func(t testing.TB, got, want int) {
		t.Helper()
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	}

	for _, tt := range tt_Find_dups_in_range {
		t.Run("test find dups in range code", func(t *testing.T) {

			if tt.got != tt.want {
				t.Logf("value %q, got %d, want %d", tt.value, tt.got, tt.want)
				helperInt(t, tt.got, tt.want)
			}

		})

	}

	t.Run("test P2 with example data", func(t *testing.T) {
		got := SolveP2(example_data)
		want := 4174379265

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})


	t.Run("test P2 with my input", func(t *testing.T) {
		got := SolveP2(my_aoc_test_data)
		want := 41823587546

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}
