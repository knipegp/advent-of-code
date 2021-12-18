package day5

// SolvePart2 returns the number of overlapping lines including diagonals.
func SolvePart2(input string) (overlapCount int, err error) {
	parsedLinePoints, err := parseInput(input)
	if err == nil {
		maxX, maxY := getExtremePoint(parsedLinePoints)
		floorCoords := newFloorLines(maxX+1, maxY+1)
		for _, line := range parsedLinePoints {
			floorCoords.drawLine(line)
		}
		overlapCount = floorCoords.countGreaterThan(1)
	}
	return overlapCount, err
}
