package day_2

import (
	"aoc-2025/internal/ints"
	"regexp"
	"strconv"
)

type productIdRanges struct {
	idRanges []productIdRange
}
type productIdRange struct {
	min, max int
}

func (r *productIdRanges) iterate() <-chan int { // Returns a receive-only channel of integers
	ch := make(chan int)
	go func() {
		defer close(ch)
		for _, idRange := range r.idRanges {
			for productId := range idRange.iterate() {
				ch <- productId
			}
		}
	}()
	return ch
}

func (r *productIdRange) iterate() <-chan int { // Returns a receive-only channel of integers
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := r.min; i <= r.max; i++ {
			ch <- i
		}
	}()
	return ch
}

var tokenRegex = regexp.MustCompile("(\\d+)-(\\d+)")

func fromString(line string) productIdRanges {
	sliceOfGroups := tokenRegex.FindAllStringSubmatch(line, -1)
	idRanges := productIdRanges{
		idRanges: make([]productIdRange, len(sliceOfGroups)),
	}
	for i, groups := range sliceOfGroups {
		idRanges.idRanges[i] = fromGroups(groups)
	}
	return idRanges
}

func fromGroups(groups []string) productIdRange {
	if lower, err := strconv.Atoi(groups[1]); err != nil {
		panic(err)
	} else if upper, err := strconv.Atoi(groups[2]); err != nil {
		panic(err)
	} else {
		return productIdRange{
			min: lower,
			max: upper,
		}
	}
}

func part1(lines []string) string {
	idRanges := fromString(lines[0])
	sillyIdSum := 0
	for productId := range idRanges.iterate() {
		base10Length := ints.Base10Length(productId)
		leftDigits := ints.Isolate(productId, base10Length/2, base10Length)
		rightDigits := ints.Isolate(productId, 0, base10Length/2)
		asString := strconv.Itoa(productId)
		if len(asString)%2 == 0 && leftDigits == rightDigits {
			sillyIdSum += productId
		}
	}
	return strconv.Itoa(sillyIdSum)
}

func part2(lines []string) string {
	idRanges := fromString(lines[0])
	sillySum := 0
	for productId := range idRanges.iterate() {
		base10Length := ints.Base10Length(productId)
		factors := ints.Factors(base10Length)
		for _, factor := range factors {
			compared := ints.Isolate(productId, 0, factor)
			same := true
			for i := factor; i < base10Length; i += factor {
				other := ints.Isolate(productId, i, i+factor)
				if compared != other {
					same = false
					break
				}
			}
			if same {
				sillySum += productId
				break
			}
		}
	}
	return strconv.Itoa(sillySum)
}
