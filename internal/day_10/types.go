package day_10

import (
	"container/heap"
	"fmt"
	"math"
	"slices"
)

type machine struct {
	display                []bool
	buttons                [][]int
	joltages               []int
	cachedRatios           []float64
	cachedJoltsNeeded      float64
	cachedHighJoltageIndex int
}

type healthyState struct {
	health float64
	state  []int
}
type stateHeap []healthyState

func (h stateHeap) Len() int { return len(h) }
func (h stateHeap) Less(i, j int) bool {
	return h[i].health > h[j].health
}
func (h stateHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *stateHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(healthyState))
}

func (h *stateHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (m *machine) findMinButtons() int64 {
	startingDisplay := make([]bool, len(m.display))
	var buttonPressCount int64
	states := [][]bool{
		startingDisplay,
	}
	for {
		if len(states) == 0 {
			panic("no path to unlock")
		}
		newStates := make([][]bool, 0, len(states)*len(m.buttons))
		for _, state := range states {
			if slices.Equal(startingDisplay, state) && buttonPressCount > 0 {
				continue
			}
			if slices.Equal(m.display, state) {
				return buttonPressCount
			}
			for _, button := range m.buttons {
				newStates = append(newStates, pressButton(state, button))
			}
		}
		states = newStates
		buttonPressCount++
	}
}

func (m *machine) findMinJoltageButtons() int64 {
	startingJoltages := make([]int, len(m.display))
	var buttonPressCount int64
	states := [][]int{
		startingJoltages,
	}
	fmt.Printf("seeking %v\n", m.joltages)
	for {
		if len(states) == 0 {
			panic("no path found")
		}
		newStates := make(stateHeap, 0, len(states)*len(m.buttons))
		heap.Init(&newStates)

		for _, state := range states {
			for _, button := range m.buttons {
				newState := pressJoltageButton(state, button)
				health := m.healthCheck(newState)
				if health == 1 {
					if slices.Equal(newState, m.joltages) {
						return buttonPressCount + 1
					}
				}
				heap.Push(&newStates, healthyState{
					health: health,
					state:  newState,
				})

			}
		}
		states = make([][]int, 0, 4000)
		for i := 0; i < 4000 && newStates.Len() > 0; i++ {
			newState := heap.Pop(&newStates).(healthyState)
			if !slices.ContainsFunc(states, func(item []int) bool {
				return slices.Equal(item, newState.state)
			}) && newState.health > 0 {
				states = append(states, newState.state)
			}
		}
		buttonPressCount++
	}
}

func totalJolts(state []int) float64 {
	var ret int
	for _, jolt := range state {
		ret += jolt
	}
	return float64(ret)
}

func (m *machine) healthCheck(potential []int) float64 {
	jolts := totalJolts(potential)
	neededJolts := m.joltsNeeded()
	joltCountRatio := jolts / neededJolts
	for i := range potential {
		if potential[i] > m.joltages[i] {
			return 0
		}
	}
	pivotIndex := m.highJoltageIndex()
	want := m.ratios()
	have := ratios(potential, pivotIndex)
	var ret float64 = 0
	for i := range want {
		denom := max(have[i], want[i])
		numer := min(have[i], want[i])
		ret += numer / denom
	}
	haveVsWantRatio := ret / float64(len(want))
	healthRatio := .91
	ret = healthRatio*haveVsWantRatio + (1-healthRatio)*joltCountRatio
	if math.IsNaN(ret) {
		ret = 0
	}
	return ret
}

func ratios(joltages []int, highJoltageIndex int) []float64 {
	ret := make([]float64, len(joltages))
	for i := range joltages {
		if joltages[i] == 0 {
			ret[i] = 1
		} else {
			ret[i] = float64(joltages[i]) / float64(joltages[highJoltageIndex])
		}
	}
	return ret
}

func (m *machine) ratios() []float64 {
	if m.cachedRatios == nil {
		m.cachedRatios = ratios(m.joltages, m.highJoltageIndex())
	}
	return m.cachedRatios
}

func (m *machine) joltsNeeded() float64 {
	if m.cachedJoltsNeeded == 0 {
		m.cachedJoltsNeeded = totalJolts(m.joltages)
	}
	return m.cachedJoltsNeeded
}

func (m *machine) highJoltageIndex() int {
	if m.cachedHighJoltageIndex == -1 {
		m.cachedHighJoltageIndex = highJoltageIndex(m.joltages)
	}
	return m.cachedHighJoltageIndex
}

func highJoltageIndex(joltages []int) (ret int) {
	ret = 0
	for i := range joltages {
		if joltages[i] > joltages[ret] {
			ret = i
		}
	}
	return
}

func pressJoltageButton(state []int, button []int) (ret []int) {
	ret = make([]int, len(state))
	copy(ret, state)
	for _, buttonIndex := range button {
		ret[buttonIndex]++
	}
	return
}

func pressButton(state []bool, button []int) (ret []bool) {
	ret = make([]bool, len(state))
	copy(ret, state)
	for _, buttonIndex := range button {
		ret[buttonIndex] = !state[buttonIndex]
	}
	return
}
