package main

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/knipegp/advent-of-code/year_2023/day2"
)

//go:embed day2/input.txt
var input string

func main() {
	gameSet, err := day2.NewGameSet(strings.TrimSpace(input))
	if err != nil {
		panic(err)
	}
	fmt.Println(gameSet.SumValid(day2.BlockCount{"red": 12, "green": 13, "blue": 14}))
	fmt.Println(gameSet.GetPower())
}
