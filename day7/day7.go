package day7

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"

	"github.com/gsmcwhirter/advent2017/lib"
)

type Node struct {
	name          string
	weight        int
	childrenNames []string
}

type ChildWeight struct {
	name       string
	origWeight int
	weight     int
}

func (n *Node) Print() {
	fmt.Printf("name: %s, weight: %d, children: [", n.name, n.weight)
	for _, childName := range n.childrenNames {
		fmt.Printf("%s, ", childName)
	}
	fmt.Println("]")
}

func (n *Node) ChildWeights(nodesByName map[string]Node) []ChildWeight {
	weights := make([]ChildWeight, 0)

	for _, childName := range n.childrenNames {
		child := nodesByName[childName]
		weights = append(weights, ChildWeight{
			child.name,
			child.weight,
			child.TotalWeight(nodesByName),
		})
	}

	return weights
}

func (n *Node) TotalWeight(nodesByName map[string]Node) (weight int) {
	weight = n.weight

	for _, childWeight := range n.ChildWeights(nodesByName) {
		weight += childWeight.weight
	}

	return
}

func (n *Node) IsBalanced(nodesByName map[string]Node) bool {
	childWeights := n.ChildWeights(nodesByName)
	if len(childWeights) == 0 {
		return true
	}

	min := math.MaxInt64
	max := math.MinInt64

	for _, childWeight := range childWeights {
		if childWeight.weight > max {
			max = childWeight.weight
		}

		if childWeight.weight < min {
			min = childWeight.weight
		}
	}

	return max == min
}

func ParseDataString(datString string) (nodesByName map[string]Node) {
	lines := strings.Split(strings.TrimSpace(datString), "\n")

	nodesByName = map[string]Node{}

	for _, line := range lines {
		r, _ := regexp.Compile(`\s*([^(\s]*)\s\((\d+)\)(?:\s*->\s*(.*))?`)
		m := r.FindStringSubmatch(line)
		// fmt.Printf("line='%s', name='%s', weight='%s', children='%s'\n", m[0], m[1], m[2], m[3])

		wght, _ := strconv.Atoi(m[2])

		node := Node{
			name:          m[1],
			weight:        wght,
			childrenNames: make([]string, 0),
		}

		if m[3] != "" {
			children := strings.Split(strings.TrimSpace(m[3]), ",")
			for _, childName := range children {
				childName = strings.TrimSpace(childName)
				node.childrenNames = append(node.childrenNames, childName)
			}
		}

		nodesByName[m[1]] = node
	}

	return
}

//LoadData stuff
func LoadData(filename string) map[string]Node {
	dat := lib.ReadFileData(filename)

	return ParseDataString(string(dat))
}

//RunPartA is a "main"
func RunPartA(filename string) {
	nodesByName := LoadData(filename)

	for key, node := range nodesByName {
		fmt.Printf("%s -> ", key)
		node.Print()
	}

	nodesAsChildren := map[string]bool{}

	for _, node := range nodesByName {
		_, exists := nodesAsChildren[node.name]
		if !exists {
			nodesAsChildren[node.name] = false
		}

		for _, childName := range node.childrenNames {
			nodesAsChildren[childName] = true
		}
	}

	for name, isChild := range nodesAsChildren {
		if !isChild {
			fmt.Println(name)
		}
	}
}

//RunPartB is a "main"
func RunPartB(filename string) {
	nodesByName := LoadData(filename)

	for key, node := range nodesByName {
		if !node.IsBalanced(nodesByName) {
			fmt.Printf("%s -> ", key)
			fmt.Println(node.ChildWeights(nodesByName))
		}
	}
}
