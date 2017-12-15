package day11

import (
	"fmt"
	"strings"

	"github.com/gsmcwhirter/advent2017/lib"
)

type Direction int
const (
	N Direction = iota
	NE
	NW
	S
	SE
	SW
	UNKNOWN_DIR
)

func DirectionDxDy(dir Direction) (dx int, dy int) {
	switch dir {
	case N:
		dx = 0
		dy = 1
		return
	case NE:
		dx = 1
		dy = 0
		return
	case NW:
		dx = -1
		dy = 1
		return
	case S:
		dx = 0
		dy = -1
		return
	case SE:
		dx = 1
		dy = -1
		return
	case SW:
		dx = -1
		dy = 0
		return
	default:
		dx = 0
		dy = 0
		return
	}
}

func StringToDirection(s string) Direction {
	if s == "s" {
		return S
	}

	if s == "se" {
		return SE
	}

	if s == "sw" {
		return SW
	}

	if s == "n" {
		return N
	}

	if s == "ne" {
		return NE
	}

	if s == "nw" {
		return NW
	}

	return UNKNOWN_DIR

}

func StepsTo(x, y int) int {

	xs := x
	ys := -y

	if xs < 0 && ys < 0 {
		xs = -xs
		ys = -ys
	}

	m := lib.IntMin(xs, ys)

	fmt.Printf("xs=%d, ys=%d, m=%d\n", xs, ys, m)

	if m <= 0 {
		return lib.IntAbs(xs) + lib.IntAbs(ys)
	}

	return m + lib.IntAbs(xs - ys)
}

func ParseDataString(datString string) []Direction {
	strDirs := strings.Split(strings.TrimSpace(datString), ",")
	dirs := make([]Direction, len(strDirs))

	for i, strDir := range strDirs {
		dirs[i] = StringToDirection(strDir)
	}

	return dirs
}

//LoadData stuff
func LoadData(filename string) []Direction {
	dat := lib.ReadFileData(filename)

	return ParseDataString(string(dat))
}

//RunPartA is a "main"
func RunPartA(filename string) {
	dirs := LoadData(filename)

	x, y := 0, 0

	for _, d := range dirs {
		dx, dy := DirectionDxDy(d)
		x += dx
		y += dy
	}

	fmt.Printf("x=%d, y=%d\n", x, y)
	fmt.Println(StepsTo(x, y))
}

//RunPartB is a "main"
func RunPartB(filename string) {
	dirs := LoadData(filename)

	x, y := 0, 0
	maxSteps := 0

	for _, d := range dirs {
		dx, dy := DirectionDxDy(d)
		x += dx
		y += dy

		steps := StepsTo(x, y)

		if steps > maxSteps {
			maxSteps = steps
		}
	}

	fmt.Println(maxSteps)
}
