package day2

import (
	"fmt"
	"strconv"
	"strings"
)

func SolveP1(s []string) int {
	res := 0
	for _, vr := range s {
		split := strings.Split(vr, "-")
		res += FindDupsInRange(split[0], split[1])
	}
	return res
}

func SolveP2(s []string) int {
	res := 0
	for _, vr := range s {
		split := strings.Split(vr, "-")
		res += FindDupsInRangeP2(split[0], split[1])
	}
	return res
}

// Finds the duplicates in the range between start and end
// A duplicate is a repeating character integer, some examples would be
// 11, 22, 33, 44, 123123, 12341234
// this will return an integer, of the sum of these duplicates
func FindDupsInRange(start, end string) int {
	res := 0
	startI, err_start := strconv.Atoi(start)
	endI, err_end := strconv.Atoi(end)
	if err_start != nil {
		fmt.Sprintf("cannot convert string %s to int", start)
	}
	if err_end != nil {
		fmt.Sprintf("cannot convert string %s to int", start)
	}

	for i := startI; i <= endI; i++ {
		str := strconv.Itoa(i)
		if IsDuplicate(str) {
			res += i
		}
	}
	return res
}

func FindDupsInRangeP2(start, end string) int {
	res := 0
	startI, err_start := strconv.Atoi(start)
	endI, err_end := strconv.Atoi(end)
	if err_start != nil {
		fmt.Sprintf("cannot convert string %s to int", start)
	}
	if err_end != nil {
		fmt.Sprintf("cannot convert string %s to int", start)
	}

	for i := startI; i <= endI; i++ {
		str := strconv.Itoa(i)
		if IsDuplicateP2(str) {
			res += i
		}
	}
	return res
}

func IsDuplicate(s string) bool {
	// find the midpoint of the string using len / 2
	middle := len(s) / 2

	odd := mod(len(s), 2) != 0

	// if odd just return false
	if odd {
		return false
	} else {
		// if even check [0:middle] [middle+1, end]
		return s[0:middle] == s[middle:]
	}
}

func IsDuplicateP2(s string) bool {
	for i:=1; i<6; i++{ 
		if groupsOfK(s, i) {
			return true
		}
	}
	return false
}

func groupsOfK(s string, k int) bool {
	if k == 1 && len(s) < 2 {
		return false
	}
	if k!= 1 && len(s) < k*2 || mod(len(s), k) != 0 {
		return false
	}

	prevGroup := s[0:k]
	for i := k - 1; i < len(s); i += k {
		newGroup := s[i-(k-1) : i+1]
		if newGroup != prevGroup {
			return false
		} else {
			prevGroup = s[i-(k-1) : i+1]
		}
	}
	return true
}

func groupsOf1(s string) bool {
	if len(s) < 2 {
		return false
	}
	for i := range len(s) {
		if s[i] != s[0] {
			return false
		}
	}
	return true
}

func mod(a, b int) int {
	return (a%b + b) % b
}
