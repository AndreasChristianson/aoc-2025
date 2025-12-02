package ints

import "math"

var memos = make(map[int][]int)

func Factors(input int) []int {
	if input < 0 {
		panic("Factors called with negative input")
	}
	if input <= 1 {
		return []int{}
	}

	if memos[input] == nil {
		factors := []int{1}
		stop := int(math.Sqrt(float64(input)))
		for i := 2; i <= stop; i++ {
			if input%i == 0 {
				factors = append(factors, i)
				flip := input / i
				if flip != i {
					factors = append(factors, flip)
				}
			}
		}
		memos[input] = factors
	}
	return memos[input]
}
