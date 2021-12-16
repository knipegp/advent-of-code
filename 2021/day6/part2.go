package day6

import (
	"sort"
)

const daysToTrack2 = 256

func getMapKeys(in map[int]int) []int {
	out := make([]int, len(in))
	idx := 0
	for key := range in {
		out[idx] = key
		idx++
	}
	return out
}

func mapDaysUntilFirstSpawnToKidsBdays(daysToTrack int) map[int][]int {
	daysToDaysToKidsSpawn := map[int][]int{}
	for day := 1; day <= daysToTrack2; day++ {
		kidsSpawnedDays, _ := countFishSpawned(day, daysToTrack2, respawnRateDays)
		daysToDaysToKidsSpawn[day] = kidsSpawnedDays
	}
	return daysToDaysToKidsSpawn
}

func initDayToFishBorn(daysToTrack int) map[int]int {
	daysToSpawnCount := map[int]int{}
	for day := 1; day <= daysToTrack2; day++ {
		daysToSpawnCount[day] = 0
	}
	return daysToSpawnCount
}

func addAll(in []int, add int) []int {
	for idx := range in {
		in[idx] += add
	}
	return in
}

// SolvePart2 models the number of fish spawned after 256 days.
func SolvePart2(input string) (fishCount int, err error) {
	var fishDaysUntilSpawn []int
	fishDaysUntilSpawn, err = parseInput(input)
	if err == nil {
		sort.Ints(fishDaysUntilSpawn)
		daysToSpawnCount := initDayToFishBorn(daysToTrack2)
		daysToDaysToKidsBdays := mapDaysUntilFirstSpawnToKidsBdays(daysToTrack2)
		for _, spawnDelay := range fishDaysUntilSpawn {
			daysToSpawnCount[spawnDelay]++
		}
		dayKeys := getMapKeys(daysToSpawnCount)
		sort.Ints(dayKeys)
		for _, dayIdx := range dayKeys {
			dayFishCount := daysToSpawnCount[dayIdx]
			kidsBdays := daysToDaysToKidsBdays[dayIdx]
			newFishDaysUntilSpawn := bdaysToRespawnCountdownFrom0(
				kidsBdays,
				newBoardSpawnDays,
			)
			for _, newSpawnDelay := range newFishDaysUntilSpawn {
				if newSpawnDelay <= daysToTrack2 {
					daysToSpawnCount[newSpawnDelay] += dayFishCount
				} else {
					fishCount += dayFishCount
				}
			}
		}
		for _, count := range daysToSpawnCount {
			fishCount += count
		}
	}
	return fishCount, err
}
