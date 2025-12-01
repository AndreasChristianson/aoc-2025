package day_1

import (
	"errors"
	"regexp"
	"strconv"
)

type direction string

const (
	left  direction = "L"
	right direction = "R"
	na    direction = "NA"
)

type Safe struct {
	current    int
	zeroHits   int
	zeroPasses int
}

func (s *Safe) spin(amount int, dir direction) {
	spins := amount / 100
	s.zeroPasses += spins
	amount %= 100
	switch dir {
	case left:
		amount *= -1
	}
	newCurrent := amount + s.current
	if newCurrent < 0 {
		if s.current > 0 {
			s.zeroPasses += 1
		}
		newCurrent += 100
	} else if newCurrent > 100 {
		s.zeroPasses += 1
	}
	s.current = newCurrent % 100
	if s.current == 0 {
		s.zeroHits += 1
	}
}

func (s *Safe) applyStrings(lines []string) {
	for _, line := range lines {
		s.applyString(line)
	}
}

var lineRegex = regexp.MustCompile("([RL])(\\d+)")

func (s *Safe) applyString(line string) {
	matches := lineRegex.FindStringSubmatch(line)
	if amount, err := strconv.Atoi(matches[2]); err != nil {
		panic(err)
	} else if dir, err := parseDirection(matches[1]); err != nil {
		panic(err)
	} else {
		s.spin(amount, dir)
	}
}

func part1(lines []string) string {
	safe := defaultSafe()
	safe.applyStrings(lines)
	return strconv.Itoa(safe.zeroHits)
}

func part2(lines []string) string {
	safe := defaultSafe()
	safe.applyStrings(lines)
	return strconv.Itoa(safe.zeroPasses + safe.zeroHits)
}

func defaultSafe() *Safe {
	return &Safe{
		current: 50,
	}
}

func parseDirection(dir string) (direction, error) {
	switch dir {
	case "L":
		return left, nil
	case "R":
		return right, nil
	default:
		return na, errors.New("invalid direction: " + dir)
	}
}
