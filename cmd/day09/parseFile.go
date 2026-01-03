package day9

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	"github.com/maximillian-whelan/advent-of-code-2025/internal/file"
)

func readFile(fp string) BoundingBox {
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
	positions := make([]Position, lc)
	lh := 0

	for s.Scan() {
		line := s.Text()
		split := strings.Split(line, ",")
		y, err := strconv.ParseFloat(split[0], 64)
		if err != nil {
			panic("Unable to parse x value")
		}

		x, err := strconv.ParseFloat(split[1], 64)
		if err != nil {
			panic("Unable to parse y value")
		}

		positions[lh] = Position{y: y, x: x}
		lh++
	}
	return BoundingBox{positions}
}
