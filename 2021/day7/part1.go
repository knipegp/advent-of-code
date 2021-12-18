package day7

import (
	"math"
	"sort"
	"strconv"
	"strings"
)

func parseInput(input string) (subPositions []int, err error) {
	for _, rawPos := range strings.Split(input, ",") {
		var pos int
		pos, err = strconv.Atoi(rawPos)
		if err == nil {
			if subPositions == nil {
				subPositions = []int{pos}
			} else {
				subPositions = append(subPositions, pos)
			}
		} else {
			break
		}
	}
	return subPositions, err
}

func median(ints []int) int {
	sorted := make([]int, len(ints))
	copy(sorted, ints)
	sort.Ints(sorted)
	medianIndex := float64(len(sorted)) / 2.0
	if float64(int(medianIndex)) == medianIndex {
		return sorted[int(medianIndex)]
	}
	return (sorted[int(math.Ceil(medianIndex))] + sorted[int(math.Floor(medianIndex))]) / 2
}

// SolvePart1 calculates the minimum cost of fuel for aligning crab subs.
func SolvePart1(input string) (fuelCost int, err error) {
	var subPositions []int
	subPositions, err = parseInput(input)
	med := median(subPositions)
	for _, pos := range subPositions {
		fuelCost += int(math.Abs(float64(med) - float64(pos)))
	}
	return fuelCost, err
}
