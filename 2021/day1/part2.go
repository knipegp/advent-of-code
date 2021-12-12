package day1

func (in depthVector) takeSum() int {
	sum := 0
	for _, elem := range in {
		sum += elem
	}
	return sum
}

func (in depthVector) takeWindowSum(windowLength int) depthVector {
	windowSums := make([]int, len(in)-windowLength+2)
	for vectorIdx := range windowSums {
		windowSums[vectorIdx] = in[vectorIdx : vectorIdx+windowLength].takeSum()
	}
	return windowSums
}

// SolvePart2 solves the problem and returns the number of depth increases for
// the sonar input with windowed sums.
func SolvePart2(input string) (incs int, err error) {
	parsed, err := parseInput(input)
	if err == nil {
		incs = parsed.takeWindowSum(3).takeDerivative().countGreaterThan(0)
	}
	return incs, err
}
