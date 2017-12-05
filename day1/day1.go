package day1

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gsmcwhirter/advent2017/lib"
)

//ParseDataString splits on spaces
func ParseDataString(datString string) []int {
	chars := []rune(strings.TrimSpace(datString))

	digits := make([]int, len(chars))
	for i, byt := range chars {
		digInt, _ := strconv.Atoi(string(byt))
		digits[i] = digInt
	}

	return digits
}

//LoadData stuff
func LoadData(filename string) []int {
	dat := lib.ReadFileData(filename)

	return ParseDataString(string(dat))
}

//RunPartA is a "main"
func RunPartA(filename string) {
	intList := LoadData(filename)
	numInts := len(intList)

	sum := 0
	for i, currInt := range intList {
		iNext := lib.IntMod(i+1, numInts)
		if currInt == intList[iNext] {
			sum += currInt
		}
	}

	fmt.Printf("%d\n", sum)
}

//RunPartB is a "main"
func RunPartB(filename string) {
	intList := LoadData(filename)
	numInts := len(intList)
	skipBy := numInts / 2

	sum := 0
	for i, currInt := range intList {
		iNext := lib.IntMod(i + skipBy, numInts)
		if currInt == intList[iNext] {
			sum += currInt
		}
	}

	fmt.Printf("%d\n", sum)
}
