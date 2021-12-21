package day8

import (
	"fmt"
	"regexp"
	"strings"
)

type segment string

const (
	segmentCount1 = 2
	segmentCount4 = 4
	segmentCount7 = 3
	segmentCount8 = 7
)

func isUniqueSegmentCount(count int) bool {
	for _, testCount := range []int{segmentCount1, segmentCount4, segmentCount7, segmentCount8} {
		if count == testCount {
			return true
		}
	}
	return false
}

type displayedDigit []segment

const (
	digitsPerDisplay = 4
)

func fromString(segments string) displayedDigit {
	chars := strings.Split(segments, "")
	newDigit := make([]segment, len(chars))
	for segmentIdx, seg := range chars {
		newDigit[segmentIdx] = segment(seg)
	}
	return newDigit
}

func repeatedDigitsPattern(count int) string {
	builder := strings.Builder{}
	for idx := 0; idx < count; idx++ {
		builder.WriteString("([a-z]+)")
		if idx != count-1 {
			builder.WriteString(" ")
		}
	}
	return builder.String()
}

func getDisplayedNumbers(
	entry string,
) (digits [digitsPerDisplay]displayedDigit, errOut error) {
	displayDigitsPattern := fmt.Sprintf(
		"\\| %s",
		repeatedDigitsPattern(digitsPerDisplay),
	)
	pattern, err := regexp.Compile(displayDigitsPattern)
	if err != nil {
		errOut = fmt.Errorf(
			"could not compile regex pattern %s due to error %w",
			displayDigitsPattern,
			err,
		)
	}
	displayedSegments := pattern.FindStringSubmatch(entry)
	if len(displayedSegments) == 0 {
		errOut = fmt.Errorf("could not find displayed numbers in entry %s", entry)
	} else {
		displayedSegments = displayedSegments[1:]
		for displayIdx := 0; displayIdx < digitsPerDisplay; displayIdx++ {
			digits[displayIdx] = fromString(
				displayedSegments[displayIdx],
			)
		}
	}
	return digits, errOut
}

func parseDisplayedNumbers(
	input string,
) (displayedDigits [][digitsPerDisplay]displayedDigit, errOut error) {
	displayEntries := strings.Split(input, "\n")
	displayedDigits = make([][digitsPerDisplay]displayedDigit, len(displayEntries))
	for entryIdx, entry := range displayEntries {
		if numbers, err := getDisplayedNumbers(entry); err == nil {
			copy(displayedDigits[entryIdx][:], numbers[:])
		} else {
			errOut = err
		}
	}
	return displayedDigits, errOut
}

// SolvePart2 returns the number of times that 1, 4, 7, or 8 appears in an input.
func SolvePart1(input string) (ans int, errOut error) {
	numberDisplays, err := parseDisplayedNumbers(input)
	if err == nil {
		for _, number := range numberDisplays {
			for _, segments := range number {
				if isUniqueSegmentCount(len(segments)) {
					ans++
				}
			}
		}
	} else {
		errOut = fmt.Errorf("could not solve day 8 part 1 due to error %w", err)
	}
	return ans, errOut
}
