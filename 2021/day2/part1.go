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

func (l location) displace(d displacement) (location, error) {
	var err error
	switch d.direction {
	case forward:
		l.horizontal += d.magnitude
	case up:
		l.vertical -= d.magnitude
	case down:
		l.vertical += d.magnitude
	default:
		err = fmt.Errorf("Received unknown direction %s", d.direction)
	}
	return l, err
}

func fromString(rawDisplacement string) (displacement, error) {
	pattern := regexp.MustCompile(`(\w+) (\d)`)
	groups := pattern.FindStringSubmatch(rawDisplacement)
	magnitude, err := strconv.Atoi(groups[2])
	return displacement{
		direction: movementDirection(groups[1]),
		magnitude: magnitude,
	}, err
}

func parseMovement(rawMovements string) ([]displacement, error) {
	displacements := []displacement{}
	var err error
	for _, rawDisplacement := range strings.Split(rawMovements, "\n") {
		var disp displacement
		disp, err = fromString(rawDisplacement)
		if err != nil {
			break
		}
		displacements = append(displacements, disp)
	}
	return displacements, err
}

// SolvePart1 solves part 1 and returns the product of coordinates.
func SolvePart1(input string) (int, error) {
	movements, err := parseMovement(input)
	var currentLocation location
	if err == nil {
		currentLocation = location{0, 0}
		for _, movement := range movements {
			currentLocation, err = currentLocation.displace(movement)
			if err != nil {
				break
			}
		}
	}
	return currentLocation.horizontal * currentLocation.vertical, err
}
