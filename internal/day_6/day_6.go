package day_6

import (
	"aoc-2025/internal/grid"
	"regexp"
	"strconv"
	"strings"
)

type operator string

func (o operator) perform(left int64, right int64) int64 {
	switch o {
	case add:
		return left + right
	case mult:
		return left * right
	default:
		panic("unknown operator")
	}
}

const (
	add  operator = "+"
	mult operator = "*"
)

type problem struct {
	cephalopodArgs []int64
	humanArgs      []int64
	op             operator
}

func (p *problem) evaluate(args []int64) int64 {
	collect := args[0]
	for i := 1; i < len(args); i++ {
		collect = p.op.perform(collect, args[i])
	}
	return collect
}

func part1(lines []string) string {
	g := grid.New(lines, categorizer)
	problems := partition(g)
	var sum int64
	for _, problem := range problems {
		sum += problem.evaluate(problem.humanArgs)
	}

	return strconv.FormatInt(sum, 10)
}

func part2(lines []string) string {
	g := grid.New(lines, categorizer)
	problems := partition(g)
	var sum int64
	for _, problem := range problems {
		sum += problem.evaluate(problem.cephalopodArgs)
	}

	return strconv.FormatInt(sum, 10)
}

func partition(g *grid.Grid[string]) []*problem {
	lastBlankCol := -1
	problems := make([]*problem, 0)
	for col := 0; col < g.Width; col++ {
		if isColBlank(g, col) {
			problems = append(problems, parseFromGrid(g, lastBlankCol, col))
			lastBlankCol = col
		}
	}
	problems = append(problems, parseFromGrid(g, lastBlankCol, g.Width))
	return problems
}

func isColBlank(g *grid.Grid[string], col int) bool {
	for row := 0; row < g.Height; row++ {
		if _, ok := g.Get(row, col); ok {
			return false
		}
	}
	return true
}

func parseFromGrid(g *grid.Grid[string], leftBlankCol int, rightBlankCol int) *problem {
	ret := &problem{
		humanArgs:      make([]int64, 0),
		cephalopodArgs: make([]int64, 0),
	}
	if op, ok := g.Get(g.Height-1, leftBlankCol+1); !ok {
		panic("operator not found")
	} else {
		ret.op = operator(op)
	}
	ret.cephalopodArgs = collectCephalopodArgs(g, leftBlankCol, rightBlankCol)
	ret.humanArgs = collectHumanArgs(g, leftBlankCol, rightBlankCol)

	return ret
}

func collectHumanArgs(g *grid.Grid[string], leftBlankCol int, rightBlankCol int) []int64 {
	ret := make([]int64, 0)
	for row := 0; row < g.Height-1; row++ {
		builder := strings.Builder{}
		for col := leftBlankCol + 1; col < rightBlankCol; col++ {
			if digit, ok := g.Get(row, col); ok {
				builder.WriteString(digit)
			}
		}
		if number, err := strconv.ParseInt(builder.String(), 10, 64); err != nil {
			panic(err)
		} else {
			ret = append(ret, number)
		}
	}
	return ret
}
func collectCephalopodArgs(g *grid.Grid[string], leftBlankCol int, rightBlankCol int) []int64 {
	ret := make([]int64, 0)
	for col := leftBlankCol + 1; col < rightBlankCol; col++ {
		builder := strings.Builder{}
		for row := 0; row < g.Height-1; row++ {
			if digit, ok := g.Get(row, col); ok {
				builder.WriteString(digit)
			}
		}
		if number, err := strconv.ParseInt(builder.String(), 10, 64); err != nil {
			panic(err)
		} else {
			ret = append(ret, number)
		}
	}
	return ret
}

var charRegex = regexp.MustCompile(`[+*\d]`)

func categorizer(input int32) (string, bool) {
	if charRegex.MatchString(string(input)) {
		return string(input), true
	}
	return "", false
}
