package day2

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type movementDirection string

const (
	forward movementDirection = "forward"
	up      movementDirection = "up"
	down    movementDirection = "down"
)

type displacement struct {
	direction movementDirection
	magnitude int
}

type location struct {
	horizontal int
	vertical   int
}

func (l location) displace(d displacement) location {
	switch d.direction {
	case forward:
		l.horizontal += d.magnitude
	case up:
		l.vertical -= d.magnitude
	case down:
		l.vertical += d.magnitude
	default:
		panic(fmt.Errorf("Received unknown direction %s", d.direction))
	}
	return l
}

func fromString(rawDisplacement string) displacement {
	pattern := regexp.MustCompile(`(\w+) (\d)`)
	groups := pattern.FindStringSubmatch(rawDisplacement)
	magnitude, err := strconv.Atoi(groups[2])
	if err != nil {
		panic(err)
	}
	return displacement{direction: movementDirection(groups[1]), magnitude: magnitude}
}

func parseMovement(rawMovements string) []displacement {
	displacements := []displacement{}
	for _, rawDisplacement := range strings.Split(rawMovements, "\n") {
		displacements = append(displacements, fromString(rawDisplacement))
	}
	return displacements
}

// SolvePart1 solves part 1 and returns the product of coordinates.
func SolvePart1(input string) int {
	movements := parseMovement(input)
	currentLocation := location{0, 0}
	for _, movement := range movements {
		currentLocation = currentLocation.displace(movement)
	}
	return currentLocation.horizontal * currentLocation.vertical
}
