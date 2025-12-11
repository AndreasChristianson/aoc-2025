package day_11

import (
	"aoc-2025/internal/graph"
	"regexp"
	"strings"
)

func parseLines(lines []string) *graph.DirectedGraph[string] {
	ret := graph.NewDirectedGraph[string]()
	ret.CreateNode("out")
	for _, line := range lines {
		ret.CreateNode(parseNodeName(line))
	}

	for _, line := range lines {
		from, _ := ret.Find(parseNodeName(line))
		for _, edgeName := range parseEdges(line) {
			to, _ := ret.Find(edgeName)
			ret.CreateEdge(from, to)
		}
	}
	return ret
}

var nodeEdgesRegex = regexp.MustCompile(`^[^:]*: (.*)$`)

func parseEdges(line string) []string {
	matches := nodeEdgesRegex.FindStringSubmatch(line)
	return strings.Split(matches[1], " ")
}

var nodeNameRegex = regexp.MustCompile(`^([^:]*): .*$`)

func parseNodeName(line string) string {
	matches := nodeNameRegex.FindStringSubmatch(line)
	return matches[1]
}
