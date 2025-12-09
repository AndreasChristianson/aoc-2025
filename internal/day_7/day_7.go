package day_7

import (
	"aoc-2025/internal/graph"
	"aoc-2025/internal/grid"
	"aoc-2025/internal/int_point/int_point_2d"
	"strconv"
)

func part1(lines []string) string {
	field := grid.New(lines, spliterParser)
	g, start, _ := traverse(field)
	var splits int64
	for item := range g.Traverse(start) {
		if item.nodeType == splitter {
			splits++
		}
	}
	return strconv.FormatInt(splits, 10)
}

func traverse(field *grid.Grid[object]) (
	ret *graph.DirectedGraph[node],
	start *graph.DirectedGraphNode[node],
	end *graph.DirectedGraphNode[node],
) {
	ret = graph.NewDirectedGraph[node]()
	end = ret.CreateNode(node{
		nodeType:     final,
		nodeLocation: int_point_2d.At(field.Height, field.Width/2),
	})
	for sourceItem := range field.Find(source) {
		start = ret.CreateNode(node{
			nodeType:     source,
			nodeLocation: sourceItem.Location,
		})
		break
	}
	propagate(field, ret, end, start, start.Value.nodeLocation.Down())
	return
}

func propagate(
	field *grid.Grid[object],
	dg *graph.DirectedGraph[node],
	end *graph.DirectedGraphNode[node],
	from *graph.DirectedGraphNode[node],
	location int_point_2d.Location,
) {
	for ; ; location = location.Down() {
		if val, found := field.GetByLocation(location); !found {
			dg.CreateEdge(from, end)
			return
		} else if val == splitter {
			if directedNode, found := dg.Find(node{nodeType: splitter, nodeLocation: location}); !found {
				newNode := dg.CreateNode(node{
					nodeType:     splitter,
					nodeLocation: location,
				})
				dg.CreateEdge(from, newNode)
				propagate(field, dg, end, newNode, newNode.Value.nodeLocation.Left())
				propagate(field, dg, end, newNode, newNode.Value.nodeLocation.Right())
				return
			} else {
				dg.CreateEdge(from, directedNode)
				return
			}
		}
	}
}

type node struct {
	nodeType     object
	nodeLocation int_point_2d.Location
}
type object string

const (
	splitter object = "^"
	source   object = "S"
	empty    object = "."
	final    object = "final"
)

func spliterParser(char int32) (object, bool) {
	if char == '^' {
		return splitter, true
	}
	if char == 'S' {
		return source, true
	}
	return empty, true
}

func part2(lines []string) string {
	field := grid.New(lines, spliterParser)
	g, start, end := traverse(field)
	var splits int64
	splits = g.CountPaths(start, end)
	return strconv.FormatInt(splits, 10)
}
