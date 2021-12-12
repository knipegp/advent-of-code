package day3

import (
	"fmt"
	"math/big"
	"strings"
)

func (r reportValue) copy() (reportValue, error) {
	copied := make(reportValue, len(r))
	copiedCount := copy(copied, r)
	if copiedCount != len(r) {
		return nil, fmt.Errorf(
			"Report slice copy failed; copied %d expected %d",
			copiedCount,
			len(r),
		)
	}
	return copied, nil
}

func (r reportValue) remove(index int) reportValue {
	out := append(r[:index], r[index+1:]...)
	return out
}

// filter removes elements where shouldKeep evaluates false in place.
func filter(r reportValue, shouldKeep func(*big.Int) bool) reportValue {
	removeIdxs := []int{}
	for idx, val := range r {
		if !shouldKeep(val) {
			removeIdxs = append(removeIdxs, idx)
		}
	}
	for removedCount, removeIdx := range removeIdxs {
		r = r.remove(removeIdx - removedCount)
	}
	return r
}

func countElements(r reportValue, elem *big.Int) int {
	count := 0
	for _, reportElem := range r {
		if reportElem.Cmp(elem) == 0 {
			count++
		}
	}
	return count
}

func (d diagnosticReport) binString() string {
	builder := strings.Builder{}
	for _, val := range d.reportValue {
		format := fmt.Sprintf("%%0%db\n", d.reportBitLen)
		builder.WriteString(fmt.Sprintf(format, val))
	}
	return builder.String()
}

func filterForOxygen(
	fullReport diagnosticReport,
	bitIdx int,
) (filteredReport diagnosticReport, err error) {
	gamma := fullReport.getMostCommonBits()
	var searchReports reportValue
	if searchReports, err = fullReport.reportValue.copy(); err == nil {
		gammaBit := gamma.Bit(bitIdx)
		searchReports = filter(searchReports, func(test *big.Int) bool {
			return test.Bit(bitIdx) == gammaBit
		})
	}
	filteredReport.reportValue = searchReports
	filteredReport.reportBitLen = fullReport.reportBitLen
	return filteredReport, err
}

func findOxygenValue(
	fullReport diagnosticReport,
) (o2Level *big.Int, err error) {
	var valuesCopy reportValue
	if valuesCopy, err = fullReport.reportValue.copy(); err == nil {
		searchReport := diagnosticReport{
			reportValue:  valuesCopy,
			reportBitLen: fullReport.reportBitLen,
		}
		// Bit index is reversed in the problem to how I understood it.
		for bitIdx := fullReport.reportBitLen - 1; bitIdx >= 0; bitIdx-- {
			searchReport, err = filterForOxygen(searchReport, bitIdx)
			if err != nil {
				break
			}
		}
		if err == nil {
			if len(searchReport.reportValue) > 0 &&
				countElements(
					searchReport.reportValue,
					searchReport.reportValue[0],
				) == len(
					searchReport.reportValue,
				) {
				o2Level = searchReport.reportValue[0]
			} else {
				err = fmt.Errorf("Could not find 02 value for gamma")
			}
		}
	}
	return o2Level, err
}

func filterForScrubber(
	fullReport diagnosticReport,
	bitIdx int,
) (filteredReport diagnosticReport, err error) {
	gamma := fullReport.getMostCommonBits()
	epsilon := gammaToEpsilon(gamma, fullReport.reportBitLen)
	var searchReports reportValue
	if searchReports, err = fullReport.reportValue.copy(); err == nil {
		epsilonBit := epsilon.Bit(bitIdx)
		searchReports = filter(searchReports, func(test *big.Int) bool {
			return test.Bit(bitIdx) == epsilonBit
		})
	}
	filteredReport.reportValue = searchReports
	filteredReport.reportBitLen = fullReport.reportBitLen
	return filteredReport, err
}

func findScrubberRating(
	fullReport diagnosticReport,
) (scrubberRating *big.Int, err error) {
	var valuesCopy reportValue
	if valuesCopy, err = fullReport.reportValue.copy(); err == nil {
		searchReport := diagnosticReport{
			reportValue:  valuesCopy,
			reportBitLen: fullReport.reportBitLen,
		}
		// Bit index is reversed in the problem to how I understood it.
		for bitIdx := fullReport.reportBitLen - 1; bitIdx >= 0; bitIdx-- {
			searchReport, err = filterForScrubber(searchReport, bitIdx)
			if err != nil || len(searchReport.reportValue) == 1 {
				break
			}
		}
		if err == nil {
			if len(searchReport.reportValue) > 0 &&
				countElements(
					searchReport.reportValue,
					searchReport.reportValue[0],
				) == len(
					searchReport.reportValue,
				) {
				scrubberRating = searchReport.reportValue[0]
			} else {
				err = fmt.Errorf("Could not find Scrubber rating")
			}
		}
	}
	return scrubberRating, err
}

// SolvePart2 calculates the life support rating for the submarine.
func SolvePart2(input string) (lifeSupportRating int, err error) {
	fullReport, err := parseInput(input)
	var oxyRating, scrubberRating *big.Int
	if err == nil {
		oxyRating, err = findOxygenValue(fullReport)
	}
	if err == nil {
		scrubberRating, err = findScrubberRating(fullReport)
	}
	if err == nil {
		lifeSupportRating = int(oxyRating.Uint64() * scrubberRating.Uint64())
	}
	return lifeSupportRating, err
}
