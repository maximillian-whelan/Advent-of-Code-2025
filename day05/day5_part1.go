package day5

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type FreshPair struct {
	start int
	end   int
}

type Ingredients struct {
	fresh []FreshPair
	ids   []int
}

func SolveP1(fp string) int {
	i := ReadFreshnessFile(fp)
	return CheckForFreshness(i)
}

func CheckForFreshness(i Ingredients) int {
	sum := 0
	for _, fresh := range i.fresh {
		for idx, id := range i.ids {
			summer := fresh.end + fresh.start
			if summer-id >= fresh.start && summer-id <= fresh.end {
				sum++
				i.ids[idx] = -1
			}
		}
	}

	return sum
}

func ReadFreshnessFile(fp string) Ingredients {
	secondSection := false
	fresh := []FreshPair{}
	ids := []int{}

	f, err := os.Open(fp)
	if err != nil {
		fmt.Printf("Unable to open file %q", err.Error())
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			secondSection = true
		}

		if !secondSection {
			slice := strings.Split(line, "-")
			start, _ := strconv.Atoi(slice[0])
			end, _ := strconv.Atoi(slice[1])

			fresh = append(fresh, FreshPair{start, end})
		} else {
			id, _ := strconv.Atoi(line)
			ids = append(ids, id)
		}
	}

	return Ingredients{fresh: fresh, ids: ids}
}
