package day3

import (
	"strings"
	"strconv"
	"fmt"

	"github.com/gsmcwhirter/advent2017/lib"
	"math"
)


func ParseDataString(datString string) int {
	num, _ := strconv.Atoi(strings.TrimSpace(datString))
	return num
}

//LoadData stuff
func LoadData(filename string) int {
	dat := lib.ReadFileData(filename)

	return ParseDataString(string(dat))
}

var dirs = [4][2]int{
	{-1, 0},
	{0, -1},
	{1, 0},
	{0, 1},
}

var neighbors = [8][2]int{
	{1, 1},
	{1, 0},
	{1, -1},
	{0, 1},
	{0, -1},
	{-1, 1},
	{-1, 0},
	{-1, -1},
}

func BuildSquare(prev [][]int) (sq [][]int) {
	newSize := len(prev) + 2
	sq = make([][]int, newSize)
	for i := 0; i < newSize; i += 1 {
		sq[i] = make([]int, newSize)
	}

	for r, row := range prev {
		for c, val := range row {
			sq[r + 1][c + 1] = val
		}
	}

	r, c := newSize - 1, newSize - 1

	for leg := 0; leg < 4; leg += 1 {
		dr, dc := dirs[leg][0], dirs[leg][1]
		for step := 0; step < newSize - 1; step += 1 {
			r += dr
			c += dc
			sq[r][c] = 0
			for _, nxny := range neighbors {
				nr, nc := r + nxny[0], c + nxny[1]
				if 0 <= nr && nr < newSize && 0 <= nc && nc < newSize {
					sq[r][c] += sq[nr][nc]
				}
			}
		}
	}

	return
}

func BottomRight(sq [][]int) int {
	s := len(sq)
	return sq[s - 1][s - 1]
}

func PrintSquare(sq [][]int) {
	for _, row := range sq {
		for _, val := range row {
			fmt.Printf("%d\t", val)
		}
		fmt.Println()
	}
}

//RunPartA is a "main"
func RunPartA(filename string) {
	index := LoadData(filename)
	fmt.Printf("index: %d\n", index)
	shell := int(math.Ceil(math.Sqrt(float64(index)))) / 2
	fmt.Printf("shell: %d\n", shell)
	sideSize := 2 * shell + 1
	fmt.Printf("size: %d\n", sideSize)
	shellSq := sideSize * sideSize // bottom right corner value
	fmt.Printf("shell_sq: %d\n", shellSq)
	diff := shellSq - index
	fmt.Printf("diff: %d\n", diff)
	diffModSide := lib.IntMod(diff, sideSize - 1)
	fmt.Printf("side steps: %d\n", diffModSide)
	sideSteps := lib.IntAbs(diffModSide - (sideSize / 2))
	fmt.Printf("side steps from center: %d\n", sideSteps)
	fmt.Printf("total steps: %d\n", shell + sideSteps)
}

//RunPartB is a "main"
func RunPartB(filename string) {
	target := LoadData(filename)
	fmt.Printf("target: %d\n", target)
	start := [][]int{{1}}

	var sq [][]int
	for sq = start; BottomRight(sq) < target; sq = BuildSquare(sq) {
		fmt.Printf("br: %d\n", BottomRight(sq))
		PrintSquare(sq)
		fmt.Print("\n")
	}

	fmt.Printf("br: %d\n", BottomRight(sq))
	PrintSquare(sq)
	fmt.Print("\n")
}
