package day2

import "fmt"

type locationWithAim struct {
	location
	aim int
}

func (l locationWithAim) displace(d displacement) locationWithAim {
	switch d.direction {
	case forward:
		l.horizontal += d.magnitude
		l.vertical += d.magnitude * l.aim
	case up:
		l.aim -= d.magnitude
	case down:
		l.aim += d.magnitude
	default:
		panic(fmt.Errorf("Received unknown direction %s", d.direction))
	}
	return l
}

// SolvePart2 solves part 2 and returns the product of coordinates.
func SolvePart2(input string) int {
	movements := parseMovement(input)
	currentLocation := locationWithAim{location{0, 0}, 0}
	for _, movement := range movements {
		currentLocation = currentLocation.displace(movement)
	}
	return currentLocation.horizontal * currentLocation.vertical
}
