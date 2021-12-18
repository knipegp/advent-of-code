package day7

import (
	"math"
	"sort"
)

func weightedFuelCost(dist int) int {
	cost := 0
	for x := 1; x <= dist; x++ {
		cost += x
	}
	return cost
}

// SolvePart2 calculates the submarine fuel usage with the updated costs.
func SolvePart2(input string) (fuelCost int, err error) {
	var subPositions []int
	subPositions, err = parseInput(input)
	sort.Ints(subPositions)
	fuelCost = math.MaxInt
	for testPos := 0; testPos < subPositions[len(subPositions)-1]; testPos++ {
		testCost := 0
		for _, pos := range subPositions {
			posCost := weightedFuelCost(int(math.Abs(float64(pos) - float64(testPos))))
			testCost += posCost
		}
		if testCost < fuelCost {
			fuelCost = testCost
		}
	}
	return fuelCost, err
}
