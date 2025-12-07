package day3

import (
	"strconv"
)

func SolveP1(s []string) int {
	res := 0
	for _, battery := range s {
		joltage := joltageCalc(battery)
		println(joltage)
		res += joltage
	}
	return res
}

func joltageCalc(battery string) int {
	left := 0
	right := len(battery) - 1
	largest := 0

	TwoPointer(battery, &left, &right, &largest)

	if right < len(battery)-1 {
		left = right
		right = len(battery) - 1
	}

	TwoPointer(battery, &left, &right, &largest)

	return largest
}

func TwoPointer(battery string, left, right, largest *int) {
	for *left < *right {
		l, _ := strconv.Atoi(string(battery[*left]))
		r, _ := strconv.Atoi(string(battery[*right]))

		sum := l*10 + r
		if sum > *largest {
			*largest = sum
		}

		if l <= r {
			*left++
		} else {
			*right--
		}
	}
}
