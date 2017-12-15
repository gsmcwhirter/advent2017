package day8

import (
"fmt"
"math"
"regexp"
"strconv"
"strings"

"github.com/gsmcwhirter/advent2017/lib"
)

type Condition int
const (
	GT Condition = iota
	GTE
	LT
	LTE
	EQ
	NE
	UNKNOWN_COND
)

func StringToCondition(s string) Condition {
	if s == ">" {
		return GT
	}

	if s == ">=" {
		return GTE
	}

	if s == "<" {
		return LT
	}

	if s == "<=" {
		return LTE
	}

	if s == "==" {
		return EQ
	}

	if s == "!=" {
		return NE
	}

	return UNKNOWN_COND
}

func ConditionToString(c Condition) string {
	switch c {
	case GT:
		return ">"
	case GTE:
		return ">="
	case LT:
		return "<"
	case LTE:
		return "<="
	case EQ:
		return "=="
	case NE:
		return "!="
	default:
		return "?"
	}
}

type Operator int
const (
	INCR Operator = iota
	DECR
	UNKNOWN_OP
)

func StringToOperator(s string) Operator {
	if s == "inc" {
		return INCR
	}

	if s == "dec" {
		return DECR
	}

	return UNKNOWN_OP
}

func OperatorToString(op Operator) string {
	switch op {
	case INCR:
		return "inc"
	case DECR:
		return "dec"
	default:
		return "?"
	}
}

type CondExpression struct {
	RegisterName string
	Cond Condition
	Val int
}

func (c *CondExpression) ToString() string {
	return fmt.Sprintf("%s %s %d", c.RegisterName, ConditionToString(c.Cond), c.Val)
}

type Instruction struct {
	RegisterName string
	Op Operator
	Val int
	Cond CondExpression
}

func (i *Instruction) ToString() string {
	return fmt.Sprintf("%s %s %d if %s", i.RegisterName, OperatorToString(i.Op), i.Val, i.Cond.ToString())
}


type Register struct {
	Val int
}

func (r *Register) Incr(amt int) {
	r.Val += amt
}

func (r *Register) Decr(amt int) {
	r.Val -= amt
}

type RegisterSet struct {
	Registers map[string]Register
}

func (rs *RegisterSet) MaxRegister() (maxName string, maxVal int) {
	maxName = ""
	maxVal = math.MinInt64
	for rname, register := range rs.Registers {
		if register.Val > maxVal {
			maxVal = register.Val
			maxName = rname
		}
	}

	return
}

func (rs *RegisterSet) CondTrue(cond CondExpression) bool {
	register, found := rs.Registers[cond.RegisterName]

	var val int
	if found {
		val = register.Val
	} else {
		val = 0
	}

	switch cond.Cond {
	case GT:
		return val > cond.Val
	case GTE:
		return val >= cond.Val
	case LT:
		return val < cond.Val
	case LTE:
		return val <= cond.Val
	case EQ:
		return val == cond.Val
	case NE:
		return val != cond.Val
	default:
		return false
	}
}

func (rs *RegisterSet) Apply(inst Instruction) int {
	fmt.Print(inst.ToString())
	if !rs.CondTrue(inst.Cond) {
		fmt.Println(" [false]")
		_, maxVal := rs.MaxRegister()
		return maxVal
	} else {
		fmt.Println(" [true]")
	}

	_, found := rs.Registers[inst.RegisterName]

	if !found {
		rs.Registers[inst.RegisterName] = Register{0}
	}

	register, _ := rs.Registers[inst.RegisterName]

	switch inst.Op {
	case INCR:
		register.Incr(inst.Val)
	case DECR:
		register.Decr(inst.Val)
	}

	rs.Registers[inst.RegisterName] = register

	_, maxVal := rs.MaxRegister()
	return maxVal
}

func (rs *RegisterSet) Print() {
	for rname, r := range rs.Registers {
		fmt.Printf("%s: %d\n", rname, r.Val)
	}
	fmt.Println()
}

func ParseDataString(datString string) (instructions []Instruction) {
	lines := strings.Split(strings.TrimSpace(datString), "\n")
	instructions = make([]Instruction, len(lines))

	r, _ := regexp.Compile(`(\S+)\s+(\S+)\s+(\S+)\s+if\s+(\S+)\s+(\S+)\s+(\S+)$`)
	for i, line := range lines {
		m := r.FindStringSubmatch(line)
		regName := m[1]
		operName := m[2]
		operVal, _ := strconv.Atoi(m[3])
		condName := m[4]
		condType := m[5]
		condVal, _ := strconv.Atoi(m[6])

		instructions[i] = Instruction{
			regName,
			StringToOperator(operName),
			operVal,
			CondExpression{
				condName,
				StringToCondition(condType),
				condVal,
			},
			}
	}

	return
}

//LoadData stuff
func LoadData(filename string) []Instruction {
	dat := lib.ReadFileData(filename)

	return ParseDataString(string(dat))
}

//RunPartA is a "main"
func RunPartA(filename string) {
	rs := RegisterSet{map[string]Register{}}
	instructions := LoadData(filename)

	for _, inst := range instructions {
		rs.Apply(inst)
		rs.Print()
	}

	fmt.Println(rs.MaxRegister())
}

//RunPartB is a "main"
func RunPartB(filename string) {
	rs := RegisterSet{map[string]Register{}}
	instructions := LoadData(filename)

	maxVal := math.MinInt64
	for _, inst := range instructions {
		val := rs.Apply(inst)
		if val > maxVal {
			maxVal = val
		}
		rs.Print()
	}

	fmt.Println(maxVal)
}
