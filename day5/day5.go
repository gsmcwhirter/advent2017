package day5

import (
	"strings"
	"strconv"
	"fmt"

	"github.com/gsmcwhirter/advent2017/lib"
)


type Maze struct {
	L []int
}


func ParseDataString(datString string) Maze {
	lines := strings.Split(strings.TrimSpace(datString), "\n")

	maze := Maze{}
	maze.L = make([]int, len(lines))
	for i, line := range lines {
		v, _ := strconv.Atoi(line)
		maze.L[i] = v
	}

	return maze
}

//LoadData stuff
func LoadData(filename string) Maze {
	dat := lib.ReadFileData(filename)

	return ParseDataString(string(dat))
}

func (m *Maze) Step(from int) (to int) {
	steps := m.L[from]
	to = from + steps
	m.L[from] += 1

	return
}

func (m *Maze) Step2(from int) (to int) {
	steps := m.L[from]
	to = from + steps
	if steps >= 3 {
		m.L[from] -= 1
	} else {
		m.L[from] += 1
	}

	return
}


//RunPartA is a "main"
func RunPartA(filename string) {
	maze := LoadData(filename)
	steps := 0
	for from := 0; 0 <= from && from < len(maze.L); from = maze.Step(from) {
		steps += 1
	}
	fmt.Println(steps)
}

//RunPartB is a "main"
func RunPartB(filename string) {
	maze := LoadData(filename)
	steps := 0
	for from := 0; 0 <= from && from < len(maze.L); from = maze.Step2(from) {
		steps += 1
	}
	fmt.Println(steps)
}
