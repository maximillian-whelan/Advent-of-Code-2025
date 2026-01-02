package day6

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func SolveCephalopodHomeworkSheet(fp string) int {
	sum := 0
	problems := ReadCephalopodHomeworkSheet(fp)


	for i := range len(problems) {
		operator := problems[i][len(problems[i])-1]
		println(operator)

		switch operator {
		case "+":
			innerSum, err := strconv.Atoi(problems[i][0])
			if err != nil {
				msg, _ := fmt.Printf("Unable to parse string to number %q", problems[i][0])
				panic(msg)
			}
			for j := 1; j < len(problems[i])-1; j++ {
				val, err := strconv.Atoi(problems[i][j])
				if err != nil {
					msg, _ := fmt.Printf("Unable to parse string to number %q", problems[i][j])
					panic(msg)
				}
				innerSum += val
			}
			sum += innerSum
		case "*":
			innerSum, err := strconv.Atoi(problems[i][0])
			if err != nil {
				msg, _ := fmt.Printf("Unable to parse string to number %q", problems[i][0])
				panic(msg)
			}
			for j := 1; j < len(problems[i])-1; j++ {
				val, err := strconv.Atoi(problems[i][j])
				if err != nil {
					msg, _ := fmt.Printf("Unable to parse string to number %q", problems[i][j])
					panic(msg)
				}
				innerSum *= val
			}
			sum += innerSum
		}
	}

	return sum
}

func ReadCephalopodHomeworkSheet(fp string) [][]string {
	f, err := os.Open(fp)
	if err != nil {
		panicMsg := fmt.Sprintf("can't find file with this name %q", fp)
		panic(panicMsg)
	}

	defer f.Close()

	numRows, err := rowCounter(f)
	if err != nil {
		panicMsg := fmt.Sprintf("error counting lines on file: %q", fp)
		panic(panicMsg)
	}

	f.Seek(0, 0) // reset to begining
	numCols := columnCounter(f)

	data := make([][]string, numCols)
	for i := range numCols {
		data[i] = make([]string, numRows)
	}

	f.Seek(0, 0) // reset to begining
	s := bufio.NewScanner(f)
	lh := 0
	for s.Scan() {
		line := s.Text()
		values := strings.Fields(line)

		for i, val := range values {
			data[i][lh] = val
		}
		lh++
	}

	return data
}

func rowCounter(r io.Reader) (int, error) {
	buf := make([]byte, 32*1024)
	count := 0
	lineSep := []byte{'\n'}

	for {
		char, err := r.Read(buf)
		count += bytes.Count(buf[:char], lineSep)

		switch {
		case err == io.EOF:
			return count, nil
		case err != nil:
			return count, err
		}
	}
}

func columnCounter(r io.Reader) int {
	s := bufio.NewScanner(r)
	count := 0
	for s.Scan() {
		line := s.Text()
		values := strings.Fields(line)
		count = len(values)
		break
	}
	return count
}
