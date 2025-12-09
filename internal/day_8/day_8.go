package day_8

import (
	"aoc-2025/internal/int_point/int_point_3d"
	"aoc-2025/internal/strings"
	"regexp"
	"slices"
	"strconv"
)

type decorationProject struct {
	junctionBoxes []*junctionBox
	circuits      []*circuit
	tupleMap      map[junctionBoxTuple]float64
	distanceMap   map[float64][]junctionBoxTuple
	distances     []float64
}

func (p *decorationProject) populateDistances() {
	for leftIndex := 0; leftIndex < len(p.junctionBoxes)-1; leftIndex++ {
		for rightIndex := leftIndex + 1; rightIndex < len(p.junctionBoxes); rightIndex++ {
			p.findDistance(p.junctionBoxes[leftIndex], p.junctionBoxes[rightIndex])
		}
	}
	slices.Sort(p.distances)
}

func (p *decorationProject) findDistance(left *junctionBox, right *junctionBox) {
	distance := left.location.DistanceTo(right.location)
	p.distances = append(p.distances, distance)
	tuple := junctionBoxTuple{left: left, right: right}
	if tuples, found := p.distanceMap[distance]; found {
		p.distanceMap[distance] = append(tuples, tuple)
	} else {
		p.distanceMap[distance] = []junctionBoxTuple{tuple}
	}
	p.tupleMap[tuple] = distance
}

func (p *decorationProject) makeConnection(tuple junctionBoxTuple) {
	keep := tuple.right.parentCircuit
	discard := tuple.left.parentCircuit
	if keep == discard {
		return
	}
	for _, jb := range discard.junctionBoxes {
		keep.junctionBoxes = append(keep.junctionBoxes, jb)
		jb.parentCircuit = keep
	}
	p.circuits = slices.DeleteFunc(p.circuits, func(c *circuit) bool { return c == discard })
}

func (p *decorationProject) sortCircuits() {
	slices.SortFunc(p.circuits, func(left, right *circuit) int {
		return len(right.junctionBoxes) - len(left.junctionBoxes)
	})
}

type junctionBox struct {
	location      int_point_3d.Location
	parentCircuit *circuit
}

type circuit struct {
	junctionBoxes []*junctionBox
}

type junctionBoxTuple struct {
	left, right *junctionBox
}

func part1(lines []string, connectionCount int) string {
	decoratingProject := parseLines(lines)
	decoratingProject.populateDistances()
	for i := 0; i < connectionCount; i++ {
		lowestDistance := decoratingProject.distances[i]
		tuples := decoratingProject.distanceMap[lowestDistance]
		tuple := tuples[0]
		decoratingProject.distanceMap[lowestDistance] = tuples[1:]
		decoratingProject.makeConnection(tuple)
	}
	decoratingProject.sortCircuits()
	var output int64 = 1
	for i := 0; i < 3; i++ {
		output *= int64(len(decoratingProject.circuits[i].junctionBoxes))
	}
	return strconv.FormatInt(output, 10)
}

var lineRegex = regexp.MustCompile(`^(\d+),(\d+),(\d+)$`)

func parseLines(lines []string) *decorationProject {
	ret := &decorationProject{
		tupleMap:      make(map[junctionBoxTuple]float64),
		distanceMap:   make(map[float64][]junctionBoxTuple),
		distances:     make([]float64, 0),
		circuits:      make([]*circuit, len(lines)),
		junctionBoxes: make([]*junctionBox, len(lines)),
	}
	for i, line := range lines {
		matches := lineRegex.FindStringSubmatch(line)
		newCircuit := &circuit{
			junctionBoxes: make([]*junctionBox, 1),
		}
		newJunctionBox := &junctionBox{
			parentCircuit: newCircuit,
			location: int_point_3d.At(
				strings.MustParse(matches[1]),
				strings.MustParse(matches[2]),
				strings.MustParse(matches[3]),
			),
		}
		newCircuit.junctionBoxes[0] = newJunctionBox
		ret.junctionBoxes[i] = newJunctionBox
		ret.circuits[i] = newCircuit
	}
	return ret
}

func part2(lines []string) string {
	decoratingProject := parseLines(lines)
	decoratingProject.populateDistances()
	var lastConnection junctionBoxTuple
	for i := 0; len(decoratingProject.circuits) > 1; i++ {
		lowestDistance := decoratingProject.distances[i]
		tuples := decoratingProject.distanceMap[lowestDistance]
		tuple := tuples[0]
		decoratingProject.distanceMap[lowestDistance] = tuples[1:]
		decoratingProject.makeConnection(tuple)
		lastConnection = tuple
	}
	return strconv.FormatInt(int64(lastConnection.left.location.X*lastConnection.right.location.X), 10)
}
