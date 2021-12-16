package day6

import (
	"strconv"
	"strings"
)

var (
	respawnRateDays   = 7
	newBoardSpawnDays = 9
	daysToTrack1      = 80
)

func parseInput(input string) (daysUntilSpawn []int, err error) {
	daysUntilSpawn = []int{}
	for _, rawDays := range strings.Split(input, ",") {
		var parsedDays int
		parsedDays, err = strconv.Atoi(rawDays)
		if err != nil {
			break
		}
		daysUntilSpawn = append(daysUntilSpawn, parsedDays)
	}
	return daysUntilSpawn, err
}

func countInts(vals []int, desired int) int {
	count := 0
	for _, val := range vals {
		if val == desired {
			count++
		}
	}
	return count
}

func countFishSpawned(
	daysUntilNextSpawn, daysToCount, respawnRateDays int,
) (kidsBdays []int, daysLeftOver int) {
	// ) (kidsBdays []int) {
	// -1 = daysUntilNextSpawn - daysPassed
	// day  | daysUntilNextSpawn
	// init | 1
	// 1    | 0
	// 2    | 6 + new
	if daysPassed := daysUntilNextSpawn + 1; daysPassed <= daysToCount {
		kidsBdays = []int{daysPassed}
		// 3    | 5 + new
		// 4    | 4 + new
		// 5    | 3 + new
		// 6    | 2 + new
		// 7    | 1 + new
		// 8    | 0 + new
		// 9    | 6 + new + new
		daysLeftToCount := daysToCount - daysPassed
		daysLeftOver = respawnRateDays - (daysLeftToCount % respawnRateDays) - 1
		additionalKids := daysLeftToCount / respawnRateDays
		for additionalKidsIdx := 0; additionalKidsIdx < additionalKids; additionalKidsIdx++ {
			kidsBdays = append(
				kidsBdays,
				(additionalKidsIdx+1)*respawnRateDays+daysPassed,
			)
		}
	}
	return kidsBdays, daysLeftOver
	// return kidsBdays
}

func bdaysToRespawnCountdownFrom0(bdays []int, initRespawnRate int) []int {
	for bdayIdx := range bdays {
		bdays[bdayIdx] += initRespawnRate - 1
	}
	return bdays
}

// SolvePart1 calculates the number of fish after 80 days.
func SolvePart1(input string) (fishCount int, err error) {
	var fishDaysUntilSpawn []int
	fishDaysUntilSpawn, err = parseInput(input)
	if err == nil {
		for daysIdx := 0; daysIdx < len(fishDaysUntilSpawn); daysIdx++ {
			kidsBdays, _ := countFishSpawned(
				fishDaysUntilSpawn[daysIdx],
				daysToTrack1,
				respawnRateDays,
			)
			newFishDaysUntilSpawn := bdaysToRespawnCountdownFrom0(
				kidsBdays,
				newBoardSpawnDays,
			)
			if newFishDaysUntilSpawn != nil {
				fishDaysUntilSpawn = append(
					fishDaysUntilSpawn,
					newFishDaysUntilSpawn...)
			}
		}
		fishCount = len(fishDaysUntilSpawn)
	}
	return fishCount, err
}
