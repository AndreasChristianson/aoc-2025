package ints

import "math"

func Base10Length(input int) int {
	if input < 0 {
		panic("Base10Length: negative input")
	}
	if input == 0 {
		return 1
	}
	return int(math.Ceil(math.Log10(float64(input + 1))))
}
