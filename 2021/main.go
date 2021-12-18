package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/knipegp/advent-of-code/2021/day1"
	"github.com/knipegp/advent-of-code/2021/day2"
	"github.com/knipegp/advent-of-code/2021/day3"
	"github.com/knipegp/advent-of-code/2021/day4"
	"github.com/knipegp/advent-of-code/2021/day5"
	"github.com/knipegp/advent-of-code/2021/day6"
	"github.com/knipegp/advent-of-code/2021/day7"
)

func getInputFromFile(filePath string) string {
	var file *os.File
	var err error
	if file, err = os.Open(filePath); err != nil {
		panic(err)
	}
	defer file.Close()
	var inputData []byte
	if inputData, err = io.ReadAll(file); err != nil {
		panic(err)
	}
	// Is ReadAll inserting whitespace at the end of the input or am I going
	// crazy?
	return strings.TrimSpace(string(inputData))
}

type flags struct {
	inputPath *string
	day       *int
}

func parseFlags() flags {
	defer flag.Parse()
	parsedArgs := flags{
		flag.String("input-path", "", "specify the path to the problem input"),
		flag.Int("day", 0, "specify the day to solve"),
	}
	return parsedArgs
}

func main() {
	parsedArgs := parseFlags()
	var part1, part2 int
	var errPart1, errPart2 error
	switch *parsedArgs.day {
	case 1:
		part1, errPart1 = day1.SolvePart1(getInputFromFile(*parsedArgs.inputPath))
		part2, errPart2 = day1.SolvePart2(getInputFromFile(*parsedArgs.inputPath))
	case 2:
		part1, errPart1 = day2.SolvePart1(getInputFromFile(*parsedArgs.inputPath))
		part2, errPart2 = day2.SolvePart2(getInputFromFile(*parsedArgs.inputPath))
	case 3:
		part1, errPart1 = day3.SolvePart1(getInputFromFile(*parsedArgs.inputPath))
		part2, errPart2 = day3.SolvePart2(getInputFromFile(*parsedArgs.inputPath))
	case 4:
		part1, errPart1 = day4.SolvePart1(getInputFromFile(*parsedArgs.inputPath))
		part2, errPart2 = day4.SolvePart2(getInputFromFile(*parsedArgs.inputPath))
	case 5:
		part1, errPart1 = day5.SolvePart1(getInputFromFile(*parsedArgs.inputPath))
		part2, errPart2 = day5.SolvePart2(getInputFromFile(*parsedArgs.inputPath))
	case 6:
		part1, errPart1 = day6.SolvePart1(getInputFromFile(*parsedArgs.inputPath))
		part2, errPart2 = day6.SolvePart2(getInputFromFile(*parsedArgs.inputPath))
	case 7:
		part1, errPart1 = day7.SolvePart1(getInputFromFile(*parsedArgs.inputPath))
		part2, errPart2 = day7.SolvePart2(getInputFromFile(*parsedArgs.inputPath))
	default:
		panic(fmt.Errorf("Passed invalid day %d", parsedArgs.day))
	}
	fmt.Printf(
		"Part 1: %d, Error1: %v , Part 2: %d, Error2: %v\n",
		part1,
		errPart1,
		part2,
		errPart2,
	)
}
