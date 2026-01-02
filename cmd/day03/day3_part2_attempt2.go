package day3

import (
	"strconv"
)

func SolveP2attempt2(s []string) int {
	res := 0
	for _, battery := range s {
		joltage := solution(battery)
		res += joltage
	}
	return res
}

func solution(battery string) int {
	mbattery := battery
	largest, _ := strconv.Atoi(mbattery[0:12])
	left := 0
	right := 11
	idx := 0
	for idx <= len(mbattery) {
		if right < len(mbattery) {
			right++
		}
		if len(mbattery) == 12 {
			return largest
		}

		swr := mbattery[left:idx] + mbattery[idx+1:right+1]
		if len(swr) > 12 {
			panic("should never be longer than 12")
		}
		
		newValue, _ := strconv.Atoi(swr)
		if newValue > largest {
			largest = newValue
			mbattery = mbattery[left:idx] + mbattery[idx+1:]
			idx = 0    // reset to 0 as the next value could be higher than the first
			right = 11 // reset to 11 as restarting the checks
		} else {
			if idx < right {
				right--
			}
			// just gobble at 11
			if idx == 11 {
				mbattery = mbattery[left:idx] + mbattery[idx+1:]
				idx = 10
				right = 11
	 		}
			idx++
		}
	}
	return largest
}
