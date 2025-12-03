package day_3

import (
	"math"
	"strconv"
)

type batteryBank struct {
	batteries []int
}

func (b *batteryBank) maxJoltage(size int) int64 {
	maxes := make([]int, size)
	location := -1
	for i := 0; i < size; i++ {
		start := location + 1
		stop := len(b.batteries) - size + i + 1
		location = b.findHighJoltage(start, stop)
		maxes[i] = b.batteries[location]
	}
	joltage := int64(0)
	for i := 0; i < len(maxes); i++ {
		joltage += int64(maxes[i]) * int64(math.Pow10(size-i-1))
	}
	return joltage
}

func (b *batteryBank) findHighJoltage(start int, end int) int {
	largest := -1
	largestIndex := 0
	for i := start; i < end; i++ {
		if b.batteries[i] > largest {
			largest = b.batteries[i]
			largestIndex = i
		}
	}
	return largestIndex
}

func fromString(line string) batteryBank {
	bank := batteryBank{
		batteries: make([]int, len(line)),
	}
	for i, char := range line {
		bank.batteries[i] = int(char - '0')
	}
	return bank
}

func part1(lines []string) string {
	totalJoltage := int64(0)
	for _, line := range lines {
		bank := fromString(line)
		totalJoltage += bank.maxJoltage(2)
	}

	return strconv.FormatInt(totalJoltage, 10)
}

func part2(lines []string) string {
	totalJoltage := int64(0)
	for _, line := range lines {
		bank := fromString(line)
		totalJoltage += bank.maxJoltage(12)
	}

	return strconv.FormatInt(totalJoltage, 10)
}
