package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/knipegp/advent-of-code/2021/day1"
	"github.com/knipegp/advent-of-code/2021/day2"
	"github.com/knipegp/advent-of-code/2021/day3"
	"github.com/knipegp/advent-of-code/2021/day4"
	"github.com/knipegp/advent-of-code/2021/day5"
	"github.com/knipegp/advent-of-code/2021/day6"
	"github.com/knipegp/advent-of-code/2021/day7"
	"github.com/knipegp/advent-of-code/2021/day8"
)

const (
	day1Num int = 1
	day2Num     = 2
	day3Num     = 3
	day4Num     = 4
	day5Num     = 5
	day6Num     = 6
	day7Num     = 7
	day8Num     = 8
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
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	switch *parsedArgs.day {
	case day1Num:
		part1, errPart1 = day1.SolvePart1(getInputFromFile(*parsedArgs.inputPath))
		part2, errPart2 = day1.SolvePart2(getInputFromFile(*parsedArgs.inputPath))
	case day2Num:
		part1, errPart1 = day2.SolvePart1(getInputFromFile(*parsedArgs.inputPath))
		part2, errPart2 = day2.SolvePart2(getInputFromFile(*parsedArgs.inputPath))
	case day3Num:
		part1, errPart1 = day3.SolvePart1(getInputFromFile(*parsedArgs.inputPath))
		part2, errPart2 = day3.SolvePart2(getInputFromFile(*parsedArgs.inputPath))
	case day4Num:
		part1, errPart1 = day4.SolvePart1(getInputFromFile(*parsedArgs.inputPath))
		part2, errPart2 = day4.SolvePart2(getInputFromFile(*parsedArgs.inputPath))
	case day5Num:
		part1, errPart1 = day5.SolvePart1(getInputFromFile(*parsedArgs.inputPath))
		part2, errPart2 = day5.SolvePart2(getInputFromFile(*parsedArgs.inputPath))
	case day6Num:
		part1, errPart1 = day6.SolvePart1(getInputFromFile(*parsedArgs.inputPath))
		part2, errPart2 = day6.SolvePart2(getInputFromFile(*parsedArgs.inputPath))
	case day7Num:
		part1, errPart1 = day7.SolvePart1(getInputFromFile(*parsedArgs.inputPath))
		part2, errPart2 = day7.SolvePart2(getInputFromFile(*parsedArgs.inputPath))
	case day8Num:
		part1, errPart1 = day8.SolvePart1(getInputFromFile(*parsedArgs.inputPath))
		part2, errPart2 = day8.SolvePart2(getInputFromFile(*parsedArgs.inputPath))
	default:
		panic(fmt.Errorf("passed invalid day %d", parsedArgs.day))
	}
	logger.Printf(
		"Part 1: %d, Error1: %v , Part 2: %d, Error2: %v\n",
		part1,
		errPart1,
		part2,
		errPart2,
	)
}
