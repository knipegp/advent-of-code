package day2

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
	}
	return l
}

// SolvePart2 solves part 2 and returns the product of coordinates.
func SolvePart2(input string) (prod int, err error) {
	var movements []displacement
	movements, err = parseMovement(input)
	if err == nil {
		currentLocation := locationWithAim{location{0, 0}, 0}
		for _, movement := range movements {
			currentLocation = currentLocation.displace(movement)
		}
		prod = currentLocation.horizontal * currentLocation.vertical
	}
	return prod, err
}
