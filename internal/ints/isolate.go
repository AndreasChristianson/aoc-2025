package ints

import "math"

func Isolate(target int, lowestDecimalDigit int, highestDecimalDigit int) int {
	return target % int(math.Pow10(highestDecimalDigit)) / int(math.Pow10(lowestDecimalDigit))
}
