package graph

import (
	"iter"
	"slices"
)

type DirectedGraphNode[V comparable] struct {
	Value        V
	destinations []*DirectedGraphNode[V]
}

type DirectedGraph[V comparable] struct {
	nodes          []*DirectedGraphNode[V]
	value          V
	countPathMemos map[countPathMemoKey[V]]int64
}

type countPathMemoKey[V comparable] struct {
	from, to *DirectedGraphNode[V]
}

func (n *DirectedGraph[V]) CreateNode(value V) (ret *DirectedGraphNode[V]) {
	ret = &DirectedGraphNode[V]{
		Value:        value,
		destinations: make([]*DirectedGraphNode[V], 0),
	}
	n.nodes = append(n.nodes, ret)
	return
}

func (n *DirectedGraph[V]) Traverse(start *DirectedGraphNode[V]) iter.Seq[V] {
	return func(yield func(item V) bool) {
		seen := make([]*DirectedGraphNode[V], 0)
		candidates := []*DirectedGraphNode[V]{
			start,
		}
		for len(candidates) > 0 {
			candidate := candidates[0]
			if slices.Contains(seen, candidate) {
				candidates = candidates[1:]
				continue
			}
			seen = append(seen, candidate)
			if !yield(candidate.Value) {
				return
			}
			candidates = slices.Concat(candidates[1:], candidate.destinations)
		}

	}
}

func (n *DirectedGraph[V]) CreateEdge(from *DirectedGraphNode[V], to *DirectedGraphNode[V]) {
	if !slices.Contains(n.nodes, from) {
		panic("directed graph does not contain from")
	}
	if !slices.Contains(n.nodes, to) {
		panic("directed graph does not contain to")
	}
	from.destinations = append(from.destinations, to)
}

func (n *DirectedGraph[V]) Find(value V) (*DirectedGraphNode[V], bool) {
	for _, node := range n.nodes {
		if node.Value == value {
			return node, true
		}
	}
	return nil, false
}

func (n *DirectedGraph[V]) CountPaths(
	start *DirectedGraphNode[V],
	end *DirectedGraphNode[V],
) int64 {
	memoKey := countPathMemoKey[V]{start, end}
	if memo, found := n.countPathMemos[memoKey]; found {
		return memo
	}
	var sum int64
	for _, node := range start.destinations {
		if node == end {
			sum++
		} else {
			sum += n.CountPaths(node, end)
		}
	}
	n.countPathMemos[memoKey] = sum
	return sum
}

func NewDirectedGraph[V comparable]() *DirectedGraph[V] {
	return &DirectedGraph[V]{
		nodes:          make([]*DirectedGraphNode[V], 0),
		countPathMemos: make(map[countPathMemoKey[V]]int64),
	}
}
