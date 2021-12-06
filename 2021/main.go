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
	switch *parsedArgs.day {
	case 1:
		part1 = day1.SolvePart1(getInputFromFile(*parsedArgs.inputPath))
		part2 = day1.SolvePart2(getInputFromFile(*parsedArgs.inputPath))
	case 2:
		part1 = day2.SolvePart1(getInputFromFile(*parsedArgs.inputPath))
		part2 = day2.SolvePart2(getInputFromFile(*parsedArgs.inputPath))
	case 3:
		part1 = day3.SolvePart1(getInputFromFile(*parsedArgs.inputPath))
		// part2 = day2.SolvePart2(getInputFromFile(*parsedArgs.inputPath))
	default:
		panic(fmt.Errorf("Passed invalid day %d", parsedArgs.day))
	}
	fmt.Printf("Part 1: %d, Part 2: %d\n", part1, part2)
}
