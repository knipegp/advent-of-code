package day5

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

type linePoints struct {
	start [2]int
	end   [2]int
}

type floorLines [][]int

// String prints the floor as a grid.
func (f floorLines) String() string {
	builder := strings.Builder{}
	for _, row := range f {
		builder.WriteString(fmt.Sprintf("%v\n", row))
	}
	return builder.String()
}

func (f floorLines) countGreaterThan(val int) int {
	count := 0
	for _, row := range f {
		for _, elem := range row {
			if elem > val {
				count++
			}
		}
	}
	return count
}

func greatestCommonDenominator(a, b int) int {
	smallerVal := int(math.Abs(math.Min(float64(a), float64(b))))
	if smallerVal == 0 {
		return int(math.Max(float64(a), float64(b)))
	}
	for possible := smallerVal; possible >= 1; possible-- {
		if a%possible == 0 && b%possible == 0 {
			return possible
		}
	}
	panic("Your GCD algo is bad")
}

func getExtremePoint(lines []linePoints) (maxX, maxY int) {
	var currentMaxX, currentMaxY float64
	for _, line := range lines {
		currentMaxX = math.Max(
			math.Max(float64(line.start[0]), float64(line.end[0])),
			currentMaxX,
		)
		currentMaxY = math.Max(
			math.Max(float64(line.start[1]), float64(line.end[1])),
			currentMaxY,
		)
	}
	return int(currentMaxX), int(currentMaxY)
}

func newFloorLines(xLen, yLen int) floorLines {
	newFloorLines := make([][]int, yLen)
	for rowIdx := range newFloorLines {
		newFloorLines[rowIdx] = make([]int, xLen)
	}
	return newFloorLines
}

func removeDiagonals(lines []linePoints) []linePoints {
	removeCount := 0
	filtered := make([]linePoints, len(lines))
	copy(filtered, lines)
	for lineIdx, line := range lines {
		if line.start[0] != line.end[0] && line.start[1] != line.end[1] {
			filtered = append(
				filtered[:lineIdx-removeCount],
				filtered[lineIdx-removeCount+1:]...)
			removeCount++
		}
	}
	return filtered
}

func (f floorLines) drawLine(line linePoints) {
	xDelta := line.end[0] - line.start[0]
	yDelta := line.end[1] - line.start[1]
	gcd := greatestCommonDenominator(xDelta, yDelta)
	xStep := xDelta / gcd
	yStep := yDelta / gcd
	for currentCoord := line.start; currentCoord != line.end; {
		f[currentCoord[1]][currentCoord[0]]++
		currentCoord[0] += xStep
		currentCoord[1] += yStep
	}
	f[line.end[1]][line.end[0]]++
}

func parseInput(input string) (parsedPoints []linePoints, err error) {
	linePattern := regexp.MustCompile(
		"(?P<startX>\\d+),(?P<startY>\\d+) -> (?P<endX>\\d+),(?P<endY>\\d+)",
	)
	parsedPoints = []linePoints{}
	for _, inputLine := range strings.Split(input, "\n") {
		lineMatch := linePattern.FindStringSubmatch(inputLine)
		if lineMatch == nil {
			err = fmt.Errorf("Could not parse coordinates from line %s", inputLine)
			break
		}
		newLine := linePoints{}
		if newLine.start[0], err = strconv.Atoi(
			lineMatch[linePattern.SubexpIndex("startX")],
		); err != nil {
			panic(err)
		}
		if newLine.start[1], err = strconv.Atoi(
			lineMatch[linePattern.SubexpIndex("startY")],
		); err != nil {
			panic(err)
		}
		if newLine.end[0], err = strconv.Atoi(
			lineMatch[linePattern.SubexpIndex("endX")],
		); err != nil {
			panic(err)
		}
		if newLine.end[1], err = strconv.Atoi(
			lineMatch[linePattern.SubexpIndex("endY")],
		); err != nil {
			panic(err)
		}
		parsedPoints = append(parsedPoints, newLine)
	}
	return parsedPoints, err
}

//SolvePart1 returns the total number of line overlap points.
func SolvePart1(input string) (overlapCount int, err error) {
	parsedLinePoints, err := parseInput(input)
	if err == nil {
		parsedLinePoints = removeDiagonals(parsedLinePoints)
		maxX, maxY := getExtremePoint(parsedLinePoints)
		floorCoords := newFloorLines(maxX+1, maxY+1)
		for _, line := range parsedLinePoints {
			floorCoords.drawLine(line)
		}
		overlapCount = floorCoords.countGreaterThan(1)
	}
	return overlapCount, err
}
