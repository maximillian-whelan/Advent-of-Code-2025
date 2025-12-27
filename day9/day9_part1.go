package day9

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/maximillian-whelan/aoc25/internal/file"
)

func LargestArea(fp string) int {
	positions := readFile(fp)
	largest := 0
	for i := range len(positions) {
		for j := range len(positions) {
			if i != j {
				x := math.Abs(float64(positions[i].x - positions[j].x)) + 1
				y := math.Abs(float64(positions[i].y-positions[j].y)) + 1
				area := int(x * y)

				if area > largest {
					largest = area
				}
			}
		}
	}

	return largest
}

type Positions struct {
	x int
	y int
}

func readFile(fp string) []Positions {
	fi, err := os.Open(fp)
	if err != nil {
		panic("unable to open file")
	}

	lc, err := file.CountLines(fi)
	if err != nil {
		panic(err)
	}

	fi.Seek(0, 0)

	s := bufio.NewScanner(fi)
	positions := make([]Positions, lc)
	lh := 0

	for s.Scan() {
		line := s.Text()
		split := strings.Split(line, ",")
		x, err := strconv.Atoi(split[0])
		if err != nil {
			panic("Unable to parse x value")
		}

		y, err := strconv.Atoi(split[1])
		if err != nil {
			panic("Unable to parse y value")
		}

		positions[lh] = Positions{y, x}
		lh++
	}
	return positions
}
