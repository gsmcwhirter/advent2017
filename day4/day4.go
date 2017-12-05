package day4

import (
	"strings"
	"fmt"

	"github.com/gsmcwhirter/advent2017/lib"
)


func ParseDataString(datString string) [][]string {
	lines := strings.Split(strings.TrimSpace(datString), "\n")

	passphrases := make([][]string, len(lines))
	for i, line := range lines {
		passphrases[i] = strings.Split(line, " ")
	}

	return passphrases
}

//LoadData stuff
func LoadData(filename string) [][]string {
	dat := lib.ReadFileData(filename)

	return ParseDataString(string(dat))
}

func ValidPassphrase(phrase []string) bool {
	seen := map[string]int{}

	for _, word := range phrase {
		ct := seen[word]
		if ct > 0 {
			return false
		}
		seen[word] = 1
	}

	return true
}

func WordToWordMap(word string) map[rune]int {
	m := map[rune]int{}
	for _, letter := range []rune(word) {
		v := m[letter]
		m[letter] = v + 1
	}

	return m
}

func EqualWordMaps(m1 map[rune]int, m2 map[rune]int) bool {
	for k1, v1 := range m1 {
		v2 := m2[k1]
		if v2 != v1 {
			return false
		}
	}

	for k2, v2 := range m2 {
		v1 := m1[k2]
		if v1 != v2 {
			return false
		}
	}

	return true
}

func ValidPassphrase2(phrase []string) bool {
	seenWordMaps := make([]map[rune]int, 0)
	for _, word := range phrase {
		m := WordToWordMap(word)

		for _, m2 := range seenWordMaps {
			if EqualWordMaps(m, m2) {
				return false
			}
		}

		seenWordMaps = append(seenWordMaps, m)
	}

	return true
}


//RunPartA is a "main"
func RunPartA(filename string) {
	phrases := LoadData(filename)

	okayCt := 0
	for _, phrase := range phrases {
		fmt.Println(phrase)
		if ValidPassphrase(phrase) {
			okayCt += 1
		}
	}

	fmt.Printf("%d\n", okayCt)
}

//RunPartB is a "main"
func RunPartB(filename string) {
	phrases := LoadData(filename)

	okayCt := 0
	for _, phrase := range phrases {
		fmt.Println(phrase)
		if ValidPassphrase2(phrase) {
			okayCt += 1
		}
	}

	fmt.Printf("%d\n", okayCt)
}
