package day_11

import (
	"strconv"
)

func part1(lines []string) (output string) {
	graph := parseLines(lines)
	you, _ := graph.Find("you")
	out, _ := graph.Find("out")
	count := graph.CountPaths(you, out)
	return strconv.FormatInt(count, 10)
}

func part2(lines []string) (output string) {
	graph := parseLines(lines)
	svr, _ := graph.Find("svr")
	out, _ := graph.Find("out")
	dac, _ := graph.Find("dac")
	fft, _ := graph.Find("fft")
	svrToDac := graph.CountPaths(svr, dac)
	dacToFft := graph.CountPaths(dac, fft)
	fftToOut := graph.CountPaths(fft, out)
	svrToFft := graph.CountPaths(svr, fft)
	fftToDac := graph.CountPaths(fft, dac)
	dacToOut := graph.CountPaths(dac, out)
	path1Count := svrToDac * dacToFft * fftToOut
	path2Count := svrToFft * fftToDac * dacToOut
	return strconv.FormatInt(path2Count+path1Count, 10)
}
