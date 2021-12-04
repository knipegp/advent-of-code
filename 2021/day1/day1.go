package day1

import (
	"bufio"
	"os"
	"strconv"
)

type depthVector []int

func getDepthsFromInput() depthVector {
	inputFile, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer inputFile.Close()
	inputScanner := bufio.NewScanner(inputFile)

	depths := []int{}
	for inputScanner.Scan() {
		rawDepth := inputScanner.Text()
		depth, err := strconv.Atoi(rawDepth)
		if err != nil {
			panic(err)
		}
		depths = append(depths, depth)
	}
	if err := inputScanner.Err(); err != nil {
		panic(err)
	}
	return depths
}

func (in depthVector) getDerivative() depthVector {
	derivative := make([]int, len(in)-1)
	for idx := 0; idx < len(derivative); idx++ {
		derivative[idx] = in[idx+1] - in[idx]
	}
	return derivative
}

func (in depthVector) countGreaterThan(val int) int {
	condTrueCount := 0
	for _, vecElem := range in {
		if vecElem > 0 {
			condTrueCount++
		}
	}
	return condTrueCount
}

func SolvePart1() int {
	return getDepthsFromInput().getDerivative().countGreaterThan(0)
}
