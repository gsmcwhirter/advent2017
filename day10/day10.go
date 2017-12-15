package day10

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gsmcwhirter/advent2017/lib"
)

type CircularBuffer struct {
	pos int
	vals []int
	size int
}

func (cb *CircularBuffer) ReverseLength(l int) {
	if l <= 1 {
		return
	}

	for i := 0; i < l / 2; i++ {
		// swap
		s1 := lib.IntMod(cb.pos + i, cb.size)
		s2 := lib.IntMod(cb.pos + l - i - 1, cb.size)
		cb.vals[s1], cb.vals[s2] = cb.vals[s2], cb.vals[s1]
	}
}

func (cb *CircularBuffer) Advance(l int) {
	cb.pos = lib.IntMod(cb.pos + l, cb.size)
}

func NewCircularBuffer(sz int) CircularBuffer {
	cb := CircularBuffer{
		0,
		make([]int, sz),
		sz,
	}

	for i := 0; i < sz; i++ {
		cb.vals[i] = i
	}

	return cb
}

func (cb *CircularBuffer) ProdCheck() int {
	return cb.vals[0] * cb.vals[1]
}

func ParseDataString(datString string) (cbSize int, steps []int) {
	numStrs := strings.Split(strings.TrimSpace(datString), ",")
	cbSize, _ = strconv.Atoi(numStrs[0])
	steps = make([]int, len(numStrs) - 1)

	for i := 1; i < len(numStrs); i++ {
		n, _ := strconv.Atoi(numStrs[i])
		steps[i-1] = n
	}

	return
}

//LoadData stuff
func LoadData(filename string) (int, []int) {
	dat := lib.ReadFileData(filename)

	return ParseDataString(string(dat))
}

func ParseDataString2(datString string) (cbSize int, steps []int) {
	numStrs := strings.SplitN(strings.TrimSpace(datString), ",", 2)
	cbSize, _ = strconv.Atoi(numStrs[0])

	byts := []byte(numStrs[1])

	steps = make([]int, len(byts) + 5)

	for i := 0; i < len(byts); i++ {
		steps[i] = int(byts[i])
	}

	steps[len(byts)] = 17
	steps[len(byts)+1] = 31
	steps[len(byts)+2] = 73
	steps[len(byts)+3] = 47
	steps[len(byts)+4] = 23

	return
}

//LoadData stuff
func LoadData2(filename string) (int, []int) {
	dat := lib.ReadFileData(filename)

	return ParseDataString2(string(dat))
}

//RunPartA is a "main"
func RunPartA(filename string) {
	cbSize, steps := LoadData(filename)

	cb := NewCircularBuffer(cbSize)
	skip := 0
	for _, step := range steps {
		cb.ReverseLength(step)
		cb.Advance(step + skip)
		skip++
	}

	fmt.Println(cb.ProdCheck())
}

//RunPartB is a "main"
func RunPartB(filename string) {
	cbSize, steps := LoadData2(filename)

	cb := NewCircularBuffer(cbSize)

	skip := 0
	for i := 0; i < 64; i++ {
		for _, step := range steps {
			cb.ReverseLength(step)
			cb.Advance(step + skip)
			skip++
		}
	}

	dense := make([]int, 16)

	for i := 0; i < 16; i++ {
		for j := 0; j < 16; j++ {
			if j == 0 {
				dense[i] = cb.vals[i*16 + j]
			} else {
				dense[i] = dense[i] ^ cb.vals[i*16 + j]
			}
		}
	}

	for _, val := range dense {
		fmt.Printf("%02x", val)
	}

	fmt.Println()
}
