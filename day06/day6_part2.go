package day6

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func SolveCephalopodHomeworkSheetRTL(fp string) int {
	sum := 0
	problems := ReadCephalopodHomeworkSheetRTL(fp)

	for row := range len(problems) {
		for col := range len(problems[0]) {
			fmt.Printf("%q ", problems[row][col])
		}
		fmt.Println()
	}

	for i := range len(problems) {
		operator :=
			strings.Trim(problems[i][len(problems[i])-1], " ")

		switch operator {
		case "+":
			sumSlice := make([]string, len(problems[i]))
			innerSum := 0
			for j := 0; j < len(problems[i])-1; j++ {
				for k := len(problems[i][j]) - 1; k >= 0; k-- {
					// loop through right to left and add to sum slice
					sumSlice[k] = sumSlice[k] + string(problems[i][j][k])
				}
			}
			for _, val := range sumSlice {
				iVal, _ := strconv.Atoi(strings.Trim(val, " "))
				if iVal != 0 {
					innerSum += iVal
				}
			}
			sum += innerSum
		case "*":
			sumSlice := make([]string, len(problems))
			innerSum := 1
			for j := 0; j < len(problems[i])-1; j++ {
				for k := len(problems[i][j]) - 1; k >= 0; k-- {
					// loop through right to left and add to sum slice
					sumSlice[k] = sumSlice[k] + string(problems[i][j][k])
				}
			}
			for _, val := range sumSlice {
				iVal, _ := strconv.Atoi(strings.Trim(val, " "))
				if iVal != 0 {
					innerSum *= iVal
				}
			}
			sum += innerSum
		}
	}

	return sum
}

func ReadCephalopodHomeworkSheetRTL(fp string) [][]string {
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

	f.Seek(0, 0) // reset to begining
	s := bufio.NewScanner(f)
	lineText := make([][]string, numRows)

	for i := range numRows {
		lineText[i] = make([]string, 1)
	}
	lh := 0
	for s.Scan() {
		line := s.Text()
		lineText[lh] = []string{line}
		lh++
	}
	widthColumns := makeDistanceArray(lineText[len(lineText)-1], numCols)
	data := createValueList(lineText, widthColumns, numCols, numRows)

	return data
}

// takes in the array of the strings of the operators
// returns an array of the width of the columns to take for each element
func makeDistanceArray(ops []string, numCols int) []int {
	opString := ops[0]
	cs := make([]int, numCols)
	count := 1
	col := 0
	// zeroth index value is always the first operator
	for i := 1; i < len(opString); i++ {
		if opString[i] != ' ' {
			count--
			cs[col] = count
			col++
			count = 0
		}
		count++
	}
	cs[col] = count
	return cs
}

func createValueList(lineText [][]string, widthColumns []int, numCols, numRows int) [][]string {
	data := make([][]string, numCols)
	for i := range numCols {
		data[i] = make([]string, numRows)
	}

	for i := range len(lineText) {
		for j := range len(lineText[i]) {
			line := lineText[i][j]
			marker := 0
			for k, width := range widthColumns {
				value := line[marker : marker+width]
				println(value)
				data[k][i] = value
				marker = width + marker
				marker++ // remove a space after each time
			}
		}
	}

	return data
}
