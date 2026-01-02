package day2

import (
	"fmt"
	"strconv"
	"strings"
)


func SolveP2(s []string) int {
	res := 0
	for _, vr := range s {
		split := strings.Split(vr, "-")
		res += FindDupsInRangeP2(split[0], split[1])
	}
	return res
}

func FindDupsInRangeP2(start, end string) int {
	res := 0
	startI, err_start := strconv.Atoi(start)
	endI, err_end := strconv.Atoi(end)
	if err_start != nil {
		fmt.Printf("cannot convert string %s to int", start)
	}
	if err_end != nil {
		fmt.Printf("cannot convert string %s to int", start)
	}

	for i := startI; i <= endI; i++ {
		str := strconv.Itoa(i)
		if IsDuplicateP2(str) {
			res += i
		}
	}
	return res
}

func IsDuplicateP2(s string) bool {
	for i := 1; i < 6; i++ {
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
	if k != 1 && len(s) < k*2 || mod(len(s), k) != 0 {
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
