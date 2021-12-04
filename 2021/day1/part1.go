package day1

import (
	"bufio"
	"os"
	"strconv"
)

type depthVector []int

func getDepthsFromInput(filePath string) depthVector {
	inputFile, err := os.Open(filePath)
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
func SolvePart1() int {
	return getDepthsFromInput("input.txt").takeDerivative().countGreaterThan(0)
}
