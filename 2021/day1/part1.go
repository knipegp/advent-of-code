package day1

import (
	"strconv"
	"strings"
)

type depthVector []int

func parseInput(input string) depthVector {
	depths := []int{}
	for _, rawDepth := range strings.Split(input, "\n") {
		depth, err := strconv.Atoi(rawDepth)
		if err != nil {
			panic(err)
		}
		depths = append(depths, depth)
	}
	return depths
}

func (in depthVector) takeDerivative() depthVector {
	derivative := make([]int, len(in)-1)
	for idx := 0; idx < len(derivative); idx++ {
		derivative[idx] = in[idx+1] - in[idx]
	}
	return derivative
}

func (in depthVector) countGreaterThan(val int) int {
	condTrueCount := 0
	for _, vecElem := range in {
		if vecElem > val {
			condTrueCount++
		}
	}
	return condTrueCount
}

// SolvePart1 solves the problem and returns the number of depth increases in the input file.
func SolvePart1(input string) int {
	return parseInput(input).takeDerivative().countGreaterThan(0)
}
