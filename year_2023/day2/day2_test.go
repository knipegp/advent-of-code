package day2

import (
	"reflect"
	"testing"
)

func TestCreateNewRound(t *testing.T) {
	testString := " 3 blue, 4 red"
	round, err := NewRound(testString)
	if err != nil {
		panic(err)
	}
	expectedResult := BlockCount{"blue": 3, "red": 4}
	if !reflect.DeepEqual(round, expectedResult) {
		t.Errorf("Test string %s resulted in round %v, expecting %v", testString, round, expectedResult)
	}
}

func TestCreateNewGame(t *testing.T) {
	testString := "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"
	game, err := NewGame(testString)
	if err != nil {
		panic(err)
	}
	expectedResult := Game{1, []BlockCount{{"blue": 3, "red": 4}, {"red": 1, "green": 2, "blue": 6}, {"green": 2}}}
	if !reflect.DeepEqual(game, expectedResult) {
		t.Errorf("Test string %s resulted in game %v, expecting %v", testString, game, expectedResult)
	}
}

const testGameSetString string = `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`

func TestValidSum(t *testing.T) {
	gameSet, err := NewGameSet(testGameSetString)
	if err != nil {
		panic(err)
	}
	expectedSum := 8
	if sum := gameSet.SumValid(BlockCount{"red": 12, "green": 13, "blue": 14}); sum != expectedSum {
		t.Errorf("Games resulted in sum %d, expecting %d", sum, expectedSum)
	}
}

func TestPower(t *testing.T) {
	gameSet, err := NewGameSet(testGameSetString)
	if err != nil {
		panic(err)
	}
	expectedPower := 2286
	if power := gameSet.GetPower(); power != expectedPower {
		t.Errorf("Games resulted in power %d, expecting %d", power, expectedPower)
	}
}
