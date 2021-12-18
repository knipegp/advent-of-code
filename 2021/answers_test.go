package main

import (
	"reflect"
	"runtime"
	"testing"

	"github.com/knipegp/advent-of-code/2021/data"
	"github.com/knipegp/advent-of-code/2021/day1"
	"github.com/knipegp/advent-of-code/2021/day2"
	"github.com/knipegp/advent-of-code/2021/day3"
	"github.com/knipegp/advent-of-code/2021/day4"
	"github.com/knipegp/advent-of-code/2021/day5"
	"github.com/knipegp/advent-of-code/2021/day6"
	"github.com/knipegp/advent-of-code/2021/day7"
)

type solver func(string) (int, error)

// https://stackoverflow.com/a/7053871
func getFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

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
		{day5.SolvePart1, data.Day5, 5632},
		{day5.SolvePart2, data.Day5, 22213},
		{day6.SolvePart1, data.Day6, 383160},
		{day6.SolvePart2, data.Day6, 1721148811504},
		{day7.SolvePart1, data.Day7, 343441},
		{day7.SolvePart2, data.Day7, 98925151},
	}
	for _, expected := range expectedAnswers {
		t.Run(getFunctionName(expected.getSolution), func(t *testing.T) {
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
		})
	}
}
