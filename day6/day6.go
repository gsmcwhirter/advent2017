package day6

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gsmcwhirter/advent2017/lib"
)

func ParseDataString(datString string) []int {
	chars := strings.Split(strings.TrimSpace(datString), "\t")

	ints := make([]int, len(chars))
	for i, ch := range chars {
		val, _ := strconv.Atoi(ch)
		ints[i] = val
	}

	return ints
}

//LoadData stuff
func LoadData(filename string) []int {
	dat := lib.ReadFileData(filename)

	return ParseDataString(string(dat))
}

func ConfigsEqual(a []int, b []int) bool {
	// assume len(a) == len(b)
	for i, aVal := range a {
		if b[i] != aVal {
			return false
		}
	}

	return true
}

func ConfigInHistory(a []int, history [][]int) bool {
	for _, b := range history {
		if ConfigsEqual(a, b) {
			return true
		}
	}

	return false
}

func WhereConfigInHistory(a []int, history [][]int) int {
	for i, b := range history {
		if ConfigsEqual(a, b) {
			return i
		}
	}

	return -1
}

func NextState(state []int) []int {
	PrintState(state)
	next := make([]int, len(state))
	maxI, max := lib.IntMaxList(state)
	distDiv := max / len(state)
	distMod := lib.IntMod(max, len(state))
	fmt.Printf("distDiv=%d, distMod=%d\n", distDiv, distMod)
	for i := 0; i < len(next); i++ {
		if i == maxI {
			fmt.Printf("Index %d gets %d\n", i, distDiv)
			next[i] = distDiv
		} else {
			fmt.Printf("Index %d gets +%d\n", i, distDiv)
			next[i] = state[i] + distDiv

			if lib.IntMod(i-maxI, len(state)) <= distMod {
				fmt.Printf("Index %d gets +%d\n", i, 1)
				next[i]++
			}
		}
	}
	PrintState(next)
	fmt.Println()

	return next
}

func PrintState(state []int) {
	for _, val := range state {
		fmt.Printf("%d\t", val)
	}
	fmt.Println()
}

func PrintHistory(history [][]int) {
	for _, state := range history {
		PrintState(state)
	}
}

//RunPartA is a "main"
func RunPartA(filename string) {
	start := LoadData(filename)

	history := make([][]int, 0)

	for state := start; !ConfigInHistory(state, history); state = NextState(state) {
		history = append(history, state)
	}

	// PrintHistory(history)
	fmt.Println(len(history))
}

//RunPartB is a "main"
func RunPartB(filename string) {
	start := LoadData(filename)

	history := make([][]int, 0)
	var state []int
	for state = start; WhereConfigInHistory(state, history) == -1; state = NextState(state) {
		history = append(history, state)
	}

	// PrintHistory(history)
	fmt.Println(len(history) - WhereConfigInHistory(state, history))
}
