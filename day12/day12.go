package day12

import (
	"fmt"
	"strings"

	"github.com/gsmcwhirter/advent2017/lib"
)

type VisitQueue struct {
	Seen map[string]bool
	pos int
	size int
	Queue []string
}

func (q *VisitQueue) Resize() {
	fmt.Println("Resizing")
	newQ := make([]string, len(q.Queue) * 2)

	for i := 0; i < q.size; i++ {
		newQ[i] = q.Queue[lib.IntMod(q.pos + i, len(q.Queue))]
	}

	q.Queue = newQ
	q.pos = 0

	fmt.Println("Resizing Done")
}

func (q *VisitQueue) Push(node string) {
	_, found := q.Seen[node]
	if !found {
		if q.size == len(q.Queue) {
			q.Resize()
		}

		q.Queue[lib.IntMod(q.pos + q.size, len(q.Queue))] = node
		q.size++
		q.Seen[node] = true
	}
}

func (q *VisitQueue) Pop() (val string) {
	if q.size == 0 {
		val = ""
		return
	}

	val = q.Queue[q.pos]
	q.pos = lib.IntMod(q.pos + 1, len(q.Queue))
	q.size--

	return
}

func NewVisitQueue() VisitQueue {
	return VisitQueue{
		map[string]bool{},
		0,
		0,
		make([]string, 16),
	}
}


func ParseDataString(datString string) map[string][]string {
	lines := strings.Split(strings.TrimSpace(datString), "\n")

	m := map[string][]string{}

	for _, line := range lines {
		parts := strings.Split(line, " <-> ")
		source := strings.TrimSpace(parts[0])
		dest := strings.Split(parts[1], ",")
		destList := make([]string, len(dest))
		for i, d := range dest {
			destList[i] = strings.TrimSpace(d)
		}

		m[source] = destList
	}

	return m
}

func RunVisit(m map[string][]string, start string) VisitQueue {
	vq := NewVisitQueue()
	vq.Push(start)

	for n := vq.Pop(); n != ""; n = vq.Pop() {
		adj, found := m[n]
		if found {
			for _, adjN := range adj {
				vq.Push(adjN)
			}
		}
	}

	return vq
}


//LoadData stuff
func LoadData(filename string) map[string][]string {
	dat := lib.ReadFileData(filename)

	return ParseDataString(string(dat))
}

//RunPartA is a "main"
func RunPartA(filename string) {
	m := LoadData(filename)
	vq := RunVisit(m, "0")

	fmt.Println(len(vq.Seen))
}

//RunPartB is a "main"
func RunPartB(filename string) {
	m := LoadData(filename)
	seen := map[string]bool{}

	groups := 0
	for ; len(seen) < len(m); {
		start := ""
		for s := range m {
			if !seen[s] {
				start = s
				break
			}
		}

		if start == "" {
			break
		}

		vq := RunVisit(m, start)
		for s := range vq.Seen {
			seen[s] = true
		}
		groups++
	}

	fmt.Println(groups)
}
