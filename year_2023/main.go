package main

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/knipegp/advent-of-code/year_2023/day1"
)

//go:embed day1/input.txt
var input string

func filterEmptyLines(words []string) []string {
	var filtered []string
	for _, word := range words {
		if len(strings.TrimSpace(word)) != 0 {
			filtered = append(filtered, word)
		}
	}
	return filtered
}

func main() {
	words := filterEmptyLines(strings.Split(input, "\n"))
	var total int
	var err error
	if total, err = day1.GetCalibrationTotalDigits(words); err != nil {
		panic(err)
	}
	fmt.Println(total)
	if total, err = day1.GetCalibrationTotalTokens(words); err != nil {
		panic(err)
	}
	fmt.Println(total)
}
