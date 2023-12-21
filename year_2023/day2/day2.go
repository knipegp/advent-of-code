package day2

import (
	"fmt"
	"strconv"
	"strings"
)

type BlockCount map[string]int

func NewRound(raw string) (BlockCount, error) {
	rawRound := strings.Split(raw, ",")
	newRound := BlockCount{}
	for _, rawBlockResult := range rawRound {
		countAndColor := strings.Split(strings.TrimSpace(rawBlockResult), " ")
		if len(countAndColor) != 2 {
			return nil, fmt.Errorf("Could not parse round from %s", raw)
		}
		count, err := strconv.Atoi(countAndColor[0])
		if err != nil {
			return nil, fmt.Errorf("Could not parse block result from %s", countAndColor)
		}
		newRound[countAndColor[1]] = count
	}
	return newRound, nil
}

type Game struct {
	gameId int
	rounds []BlockCount
}

func NewGame(raw string) (Game, error) {
	gameIdAndRounds := strings.Split(raw, ":")
	if len(gameIdAndRounds) != 2 {
		return Game{}, fmt.Errorf("Could not parse game from string %s", raw)
	}
	rawGameId := strings.Split(strings.TrimSpace(gameIdAndRounds[0]), " ")
	if len(rawGameId) != 2 {
		return Game{}, fmt.Errorf("Could not parse game ID from string %s", gameIdAndRounds[0])
	}
	gameId, err := strconv.Atoi(rawGameId[1])
	if err != nil {
		return Game{}, fmt.Errorf("Count not parse game ID from string %s", rawGameId[1])
	}
	rounds := []BlockCount{}
	for _, rawRound := range strings.Split(gameIdAndRounds[1], ";") {
		newRound, err := NewRound(rawRound)
		if err != nil {
			return Game{}, err
		}
		rounds = append(rounds, newRound)
	}
	return Game{gameId: gameId, rounds: rounds}, nil
}

func (g Game) checkValidity(availableBlocks BlockCount) bool {
	for _, round := range g.rounds {
		for color, count := range round {
			available, ok := availableBlocks[color]
			if !ok {
				return false
			}
			if count > available {
				return false
			}
		}
	}
	return true
}

func (g Game) GetPower() int {
	minBlocks := BlockCount{"red": 0, "green": 0, "blue": 0}
	for _, round := range g.rounds {
		for color, count := range round {
			if count > minBlocks[color] {
				minBlocks[color] = count
			}
		}
	}
	power := 1
	for _, count := range minBlocks {
		power *= count
	}
	return power
}

type GameSet []Game

func NewGameSet(raw string) (GameSet, error) {
	rawGames := strings.Split(raw, "\n")
	games := []Game{}
	for _, rawGame := range rawGames {
		newGame, err := NewGame(rawGame)
		if err != nil {
			return nil, err
		}
		games = append(games, newGame)
	}
	return games, nil
}

func (g GameSet) SumValid(availableBlocks BlockCount) int {
	sum := 0
	for _, game := range g {
		if game.checkValidity(availableBlocks) {
			sum += game.gameId
		}
	}
	return sum
}

func (g GameSet) GetPower() int {
	power := 0
	for _, game := range g {
		power += game.GetPower()
	}
	return power
}
