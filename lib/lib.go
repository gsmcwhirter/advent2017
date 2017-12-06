package lib

import (
	"io/ioutil"
	"math"
)

//Check causes a panic on errors
func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadFileData(filename string) string {
	dat, err := ioutil.ReadFile(filename)
	Check(err)

	return string(dat)
}

// IntMod is to make up for the verison of % that go ships... BOOOOOOOOOOOOO!!!!!!
func IntMod(a, b int) int {
	r := a % b
	if r < 0 {
		r += b
	}
	return r
}

func IntMax(a int, b int) int {
	if a >= b {
		return a
	}

	return b
}

func IntMaxList(l []int) (maxI int, max int) {
	max = l[0]
	maxI = 0

	for i, val := range l {
		if val > max {
			maxI = i
			max = val
		}
	}

	return
}

func IntMin(a int, b int) int {
	if a <= b {
		return a
	}

	return b
}

func UInt32Max(a uint32, b uint32) uint32 {
	if a >= b {
		return a
	}

	return b
}

func UInt32Min(a uint32, b uint32) uint32 {
	if a <= b {
		return a
	}

	return b
}

func IntAbs(val int) int {
	return int(math.Abs(float64(val)))
}

func IntClamp(val int, min int, max int) int {
	if val < min {
		return min
	}

	if val > max {
		return max
	}

	return val
}

type RuneCount struct {
	Rune  rune
	Count int
}

func NewRuneCounts(counts map[rune]int) []RuneCount {
	runeCounts := []RuneCount{}
	for key, val := range counts {
		runeCounts = append(runeCounts, RuneCount{key, val})
	}

	return runeCounts
}

type RuneCountSorter struct {
	Counts []RuneCount
}

func (rcs RuneCountSorter) Len() int {
	return len(rcs.Counts)
}

func (rcs RuneCountSorter) Less(i, j int) bool {
	rcA := rcs.Counts[i]
	rcB := rcs.Counts[j]

	if rcA.Count > rcB.Count {
		return true
	}

	if rcA.Count < rcB.Count {
		return false
	}

	if rcA.Rune < rcB.Rune {
		return true
	}

	return false
}

func (rcs RuneCountSorter) Swap(i, j int) {
	var tmp RuneCount

	tmp = rcs.Counts[i]
	rcs.Counts[i] = rcs.Counts[j]
	rcs.Counts[j] = tmp
}
