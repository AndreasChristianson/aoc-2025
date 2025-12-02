package ints

import "math"

// Isolate returns sub-ints from the base10 representation of an int
// ............↓↓      [from after the 5th digit] [to the 7th digit]
// Isolate(10987654321, 5,                         7               ) == 76
func Isolate(target int, lowestDecimalDigit int, highestDecimalDigit int) int {
	return target % int(math.Pow10(highestDecimalDigit)) / int(math.Pow10(lowestDecimalDigit))
}
