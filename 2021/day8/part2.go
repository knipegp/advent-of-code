package day8

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

const (
	segmentA segment = "a"
	segmentB segment = "b"
	segmentC segment = "c"
	segmentD segment = "d"
	segmentE segment = "e"
	segmentF segment = "f"
	segmentG segment = "g"
)

const allDigitsMax = 9

func getAllSegments() []segment {
	return []segment{
		segmentA,
		segmentB,
		segmentC,
		segmentD,
		segmentE,
		segmentF,
		segmentG,
	}
}

type displayDigits struct {
	mixedDigits       [allDigitsMax + 1]displayedDigit
	displayedSegments [digitsPerDisplay]displayedDigit
}

func removeSegment(list []segment, removeSegs ...segment) []segment {
	for _, seg := range removeSegs {
		segFound := false
		for listIdx, listSeg := range list {
			if listSeg == seg {
				list = append(list[:listIdx], list[listIdx+1:]...)
				segFound = true
				break
			}
		}
		if !segFound {
			panic(fmt.Errorf("could not find segment %s in list %v", seg, list))
		}
	}
	return list
}

func isInSegments(seg segment, segs []segment) bool {
	for _, testSeg := range segs {
		if seg == testSeg {
			return true
		}
	}
	return false
}

func getPossibleSegmentsForLen(segCount int) []segment {
	possible := []segment{}
	for _, segs := range numToCorrectSegmentMapping() {
		for _, seg := range segs {
			if segCount == len(segs) && !isInSegments(seg, possible) {
				possible = append(possible, seg)
			}
		}
	}
	sort.SliceStable(possible, func(a, b int) bool { return possible[a] < possible[b] })
	return possible
}

func getMixedDigitsMap(
	mixedDigits [allDigitsMax + 1]displayedDigit,
) [allDigitsMax + 1][]segment {
	var mappedSegments [allDigitsMax + 1][]segment
	possibleMixedToActualSegment := map[segment][]segment{}
	for _, key := range getAllSegments() {
		possibleMixedToActualSegment[key] = getAllSegments()
	}
	for _, segments := range mixedDigits {
		possibleSegs := getPossibleSegmentsForLen(len(segments))
		for _, seg := range segments {
			if len(possibleSegs) < len(possibleMixedToActualSegment[seg]) {
				possibleMixedToActualSegment[seg] = possibleSegs
			}
		}
	}
	fmt.Println(possibleMixedToActualSegment)
	return mappedSegments
}

func parseMixedDigits(
	entry string,
) (mixedDigits [allDigitsMax + 1]displayedDigit, errOut error) {
	mixedDigitPattern := fmt.Sprintf("%s |", repeatedDigitsPattern(allDigitsMax+1))
	var pattern *regexp.Regexp
	if pattern, errOut = regexp.Compile(mixedDigitPattern); errOut == nil {
		rawSegments := pattern.FindStringSubmatch(entry)
		if len(rawSegments) == 0 {
			errOut = fmt.Errorf("could not find mixed digits in entry %s", entry)
		} else {
			rawSegments = rawSegments[1:]
			for digitIdx := 0; digitIdx <= allDigitsMax; digitIdx++ {
				mixedDigits[digitIdx] = fromString(rawSegments[digitIdx])
			}
		}
	} else {
		errOut = fmt.Errorf("could not compile mixed digit regexp pattern due to error %w", errOut)
	}
	return mixedDigits, errOut
}

func parseDisplays(
	input string,
) (displayOutput []displayDigits, errOut error) {
	displayEntries := strings.Split(input, "\n")
	displayOutputs := make([]displayDigits, len(displayEntries))
	for entryIdx, entry := range displayEntries {
		var mixedDigits [allDigitsMax + 1]displayedDigit
		var err error
		mixedDigits, err = parseMixedDigits(entry)
		if err != nil {
			errOut = err
			break
		}
		if displayed, err := getDisplayedNumbers(entry); err == nil {
			displayOutputs[entryIdx] = displayDigits{
				mixedDigits:       mixedDigits,
				displayedSegments: displayed,
			}
		} else {
			break
		}
	}
	return displayOutputs, errOut
}

// func numSegmentsToPossibleSegments(num int) []segment {
// }

func numToCorrectSegmentMapping() [allDigitsMax + 1][]segment {
	return [allDigitsMax + 1][]segment{
		{segmentA, segmentB, segmentC, segmentE, segmentF, segmentG},
		{segmentC, segmentF},
		{segmentA, segmentC, segmentD, segmentE, segmentG},
		{segmentA, segmentC, segmentD, segmentF, segmentG},
		{segmentB, segmentC, segmentD, segmentF},
		{segmentA, segmentB, segmentD, segmentF, segmentG},
		{segmentA, segmentB, segmentD, segmentE, segmentF, segmentG},
		{segmentA, segmentC, segmentF},
		{segmentA, segmentB, segmentC, segmentD, segmentE, segmentF, segmentG},
		{segmentA, segmentB, segmentC, segmentD, segmentF, segmentG},
	}
}

// SolvePart2 returns the sum of all digits.
func SolvePart2(input string) (ans int, errOut error) {
	if displays, err := parseDisplays(input); err == nil {
		getMixedDigitsMap(displays[0].mixedDigits)
	} else {
		errOut = fmt.Errorf("could not solve day 8 part 2 due to error %w", err)
	}
	return ans, errOut
}
