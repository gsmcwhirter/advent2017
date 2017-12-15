package day9

import (
	"fmt"
	"strings"

	"github.com/gsmcwhirter/advent2017/lib"
)


//LoadData stuff
func LoadData(filename string) []rune {
	dat := lib.ReadFileData(filename)

	return []rune(strings.TrimSpace(string(dat)))
}

//RunPartA is a "main"
func RunPartA(filename string) {
	stream := LoadData(filename)

	openGroups := 0
	totalScore := 0
	openGarbage := false

	for i := 0; i < len(stream); {
		currChar := stream[i]

		if currChar == '!' {
			i += 2
			continue
		}

		if openGarbage {
			if currChar == '>' {
				openGarbage = false
			}

			i++
			continue
		}

		if !openGarbage && currChar == '<' {
			openGarbage = true
		}

		if currChar == '{' {
			openGroups += 1
		}

		if currChar == '}' {
			totalScore += openGroups
			openGroups--
		}

		i++
	}

	fmt.Println(totalScore)
}

//RunPartB is a "main"
func RunPartB(filename string) {
	stream := LoadData(filename)

	totalGarbage := 0
	openGarbage := false

	for i := 0; i < len(stream); {
		currChar := stream[i]

		if currChar == '!' {
			i += 2
			continue
		}

		if openGarbage {
			if currChar == '>' {
				openGarbage = false
			} else {
				totalGarbage++
			}

			i++
			continue
		}

		if !openGarbage && currChar == '<' {
			openGarbage = true
		}

		i++
	}

	fmt.Println(totalGarbage)
}
