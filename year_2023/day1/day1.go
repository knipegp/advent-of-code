package day1

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type digitToken struct {
	token string
	value int
}

var validTokens = []digitToken{
	{
		"one", 1,
	},
	{
		"two", 2,
	},
	{
		"three", 3,
	},
	{
		"four", 4,
	},
	{
		"five", 5,
	},
	{
		"six", 6,
	},
	{
		"seven", 7,
	},
	{
		"eight", 8,
	},
	{
		"nine", 9,
	},
	{
		"1", 1,
	},
	{
		"2", 2,
	},
	{
		"3", 3,
	},
	{
		"4", 4,
	},
	{
		"5", 5,
	},
	{
		"6", 6,
	},
	{
		"7", 7,
	},
	{
		"8", 8,
	},
	{
		"9", 9,
	},
}

type wordTokens struct {
	first digitToken
	last  digitToken
}

func findTokens(word string, tokens []digitToken) (wordTokens, error) {
	firstTokenIndex := math.MaxInt
	lastTokenIndex := -1
	foundTokens := wordTokens{}
	for _, token := range validTokens {
		testFirstIndex := strings.Index(word, token.token)
		if testFirstIndex != -1 && testFirstIndex < firstTokenIndex {
			firstTokenIndex = testFirstIndex
			foundTokens.first = token
		}
		testLastIndex := strings.LastIndex(word, token.token)
		if testLastIndex > lastTokenIndex {
			lastTokenIndex = testLastIndex
			foundTokens.last = token
		}
	}
	if firstTokenIndex == math.MaxInt || lastTokenIndex == -1 {
		return foundTokens, fmt.Errorf("Found no tokens in word %s", word)
	}
	if foundTokens.first.value == 0 || foundTokens.last.value == 0 {
		return foundTokens, fmt.Errorf("Bad found tokens %v", foundTokens)
	}
	fmt.Printf("word %s first index %d last index %d tokens %v\n", word, firstTokenIndex, lastTokenIndex, foundTokens)
	return foundTokens, nil
}

func getCalibrationValue(word string) (int, error) {
	var calibrationValue string
	noValue := true
	for _, character := range word {
		if _, err := strconv.Atoi(string(character)); err == nil {
			calibrationValue = string(character)
			noValue = false
			break
		}
	}
	if noValue {
		return 0, fmt.Errorf("Found no digits in word %s", word)
	}
	for index := len(word) - 1; index >= 0; index-- {
		if _, err := strconv.Atoi(string(word[index])); err == nil {
			calibrationValue += string(word[index])
			break
		}
	}
	value, err := strconv.Atoi(calibrationValue)
	return value, err
}

func GetCalibrationTotalDigits(values []string) (int, error) {
	total := 0
	for _, calibrationValue := range values {
		var err error
		var convertedValue int
		if convertedValue, err = getCalibrationValue(calibrationValue); err != nil {
			return 0, err
		}
		total += convertedValue
	}
	return total, nil
}

func GetCalibrationTotalTokens(values []string) (int, error) {
	total := 0
	for _, calibrationValue := range values {
		var err error
		var tokens wordTokens
		if tokens, err = findTokens(calibrationValue, validTokens); err != nil {
			return 0, err
		}
		var tokenValue int
		if tokenValue, err = strconv.Atoi(fmt.Sprint(tokens.first.value) + fmt.Sprint(tokens.last.value)); err != nil {
			panic(err)
		}
		total = total + tokenValue
	}
	return total, nil
}
