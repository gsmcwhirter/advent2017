package day2

import (
	"strings"
	"strconv"
	"fmt"

	"github.com/gsmcwhirter/advent2017/lib"
	"math"
)


func ParseDataString(datString string) [][]int {
	lines := strings.Split(strings.TrimSpace(datString), "\n")
	result := make([][]int, len(lines))

	for row, line := range lines {
		fields := strings.Split(line, "\t")
		result[row] = make([]int, len(fields))

		for col, numStr := range fields {
			num, _ := strconv.Atoi(string(numStr))
			result[row][col] = num
		}
	}

	return result
}

//LoadData stuff
func LoadData(filename string) [][]int {
	dat := lib.ReadFileData(filename)

	return ParseDataString(string(dat))
}

func LineMinMax(line []int) (min int, max int) {
	min = math.MaxInt64
	max = math.MinInt64
	for _, val := range line {
		min = lib.IntMin(val, min)
		max = lib.IntMax(val, max)
	}

	return
}

func LineDiff(line []int) (diff int) {
	min, max := LineMinMax(line)
	diff = max - min
	return
}

func LineDivisors(line []int) (bigger int, smaller int) {
	for i, val1 := range line {
		for _, val2 := range line[i+1:] {
			if lib.IntMod(val1, val2) == 0 {
				bigger = val1
				smaller = val2
				return
			}
			if lib.IntMod(val2, val1) == 0 {
				bigger = val2
				smaller = val1
				return
			}
		}
	}

	bigger = 0
	smaller = 0
	return
}

func LineQuotient(line []int) (quot int) {
	bigger, smaller := LineDivisors(line)
	if smaller == 0 {
		fmt.Printf("Error: %v\n", line)
		quot = 0
		return
	}
	quot = bigger / smaller
	return
}

//RunPartA is a "main"
func RunPartA(filename string) {
	sheet := LoadData(filename)
	checksum := 0
	for _, line := range sheet {
		checksum += LineDiff(line)
	}

	fmt.Printf("%d\n", checksum)
}

//RunPartB is a "main"
func RunPartB(filename string) {
	sheet := LoadData(filename)
	checksum := 0
	for _, line := range sheet {
		checksum += LineQuotient(line)
	}

	fmt.Printf("%d\n", checksum)
}
