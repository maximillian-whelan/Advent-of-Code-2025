package day01

import (
	"strconv"
)

func SolveP1(rots []string) int {
	turnValue := 50
	result := 0

	for _, rotations := range rots {

		direction := string(rotations[0])
		change, err := strconv.Atoi(rotations[1:])

		if err == nil {
			if direction == "R" {
				turnValue = mod(turnValue+change, 100)

				if turnValue == 0 {
					result++
				}
			}

			if direction == "L" {
				turnValue = mod(turnValue-change, 100)

				if turnValue == 0 {
					result++
				}
			}
		}
	}
	return result
}

func SolveP2(rots []string) int {
	cur := 50
	res := 0

	for _, rotations := range rots {

		direction := string(rotations[0])
		change, err := strconv.Atoi(rotations[1:])
		if err == nil {
		}

		left := direction == "L"
		if left {
			for range change {
				cur--
				if cur == -1 {
					cur = 99
				}
				if cur == 0 {
					res++
				}
			}
		} else {
			for range change {
				cur++
				if cur == 100 {
					cur = 0
				}
				if cur == 0 {
					res++
				}
			}
		}
	}
	return res
}

func mod(a, b int) int {
	return (a%b + b) % b
}
