package main

import (
	"testing"

	"github.com/knipegp/advent-of-code/2021/data"
	"github.com/knipegp/advent-of-code/2021/day1"
	"github.com/knipegp/advent-of-code/2021/day2"
	"github.com/knipegp/advent-of-code/2021/day3"
	"github.com/knipegp/advent-of-code/2021/day4"
)

type solver func(string) (int, error)

func TestAnswers(t *testing.T) {
	expectedAnswers := []struct {
		getSolution solver
		input       string
		answer      int
	}{
		{day1.SolvePart1, data.Day1, 1688},
		{day1.SolvePart2, data.Day1, 1728},
		{day2.SolvePart1, data.Day2, 1694130},
		{day2.SolvePart2, data.Day2, 1698850445},
		{day3.SolvePart1, data.Day3, 3374136},
		{day3.SolvePart2, data.Day3, 4432698},
		{day4.SolvePart1, data.Day4, 8136},
		{day4.SolvePart2, data.Day4, 12738},
	}
	for _, expected := range expectedAnswers {
		if calcSoln, err := expected.getSolution(expected.input); calcSoln != expected.answer ||
			err != nil {
			t.Errorf(
				"Solver %v returned %d, expected %d; Error %v",
				expected.getSolution,
				calcSoln,
				expected.answer,
				err,
			)
		}
	}
}
