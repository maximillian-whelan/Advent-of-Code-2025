package day5

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func SolveP2(fp string) int {

	freshness := ReadP2(fp)
	return HandleCleaning(freshness)
}

func HandleCleaning(freshness []FreshPair) int {
	// sort by start value smallest to largest
	sort.Slice(freshness, func(i, j int) bool {
		return freshness[i].start < freshness[j].start
	})

	// remove duplicates
	for i := 1; i < len(freshness); i++ {
		if freshness[i-1].start == freshness[i].start && freshness[i-1].end == freshness[i].end {
			freshness = removeIndex(freshness, i-1)
			i--
		}
	}

	j := 1
	for j < len(freshness) {
		if freshness[j-1].end >= freshness[j].start {
			if freshness[j-1].end >= freshness[j].end {
				freshness = removeIndex(freshness, j)
				if j > 1 {
					j--
				}
			} else {
				freshness[j].start = freshness[j-1].end + 1
			}
		} else {
			j++
		}
	}

	sum := 0
	for k := range freshness {
		println(freshness[k].start, freshness[k].end)
		sum += (freshness[k].end - freshness[k].start) + 1
	}

	return sum
}

func removeIndex(s []FreshPair, idx int) []FreshPair {
	return append(s[0:idx], s[idx+1:]...)
}

func ReadP2(fp string) []FreshPair {
	fresh := []FreshPair{}
	f, err := os.Open(fp)
	if err != nil {
		fmt.Printf("Unable to open file %q", err.Error())
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		slice := strings.Split(line, "-")
		start, _ := strconv.Atoi(slice[0])
		end, _ := strconv.Atoi(slice[1])
		fresh = append(fresh, FreshPair{start, end})
	}
	return fresh
}

