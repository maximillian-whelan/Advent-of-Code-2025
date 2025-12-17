package day7

import (
	"bufio"
	"fmt"
	"os"
)

func SolveTachSplitting(fp string) int {
	b := ReadManifold(fp)

	sCount := 0
	// start at row 1 find the start location and add first tach split
	for row := 1; row < len(b); row++ {
		for col := range b[row] {
			if b[row-1][col] == 'S' {
				// found start location add tach bar at row col
				b[row][col] = '|'
			}

			if b[row-1][col] == '|' {
				if b[row][col] == '^' {
					// we should split
					sCount++
					if col-1 > 0 {
						b[row][col-1] = '|'
					}
					if col+1 < len(b[row]) {
						b[row][col+1] = '|'
					}
				} else {
					// we are still falling down - add the next tach bar 
					b[row][col] = '|'
				}
			}
		}
	}

	for row := range len(b) {
		for col := range len(b[row]) {
			fmt.Printf("%q ", b[row][col])
		}
		fmt.Println()
	}
	return sCount
}

func ReadManifold(fp string) [][]byte {
	fi, err := os.Open(fp)
	println(fi.Stat())
	if err != nil {
		panicMsg, _ := fmt.Printf("Unable to open %q", fp)
		panic(panicMsg)
	}

	defer fi.Close()

	data := [][]byte{}
	s := bufio.NewScanner(fi)
	for s.Scan() {
		bytes := []byte(s.Text())
		data = append(data, bytes)
	}
	return data
}
