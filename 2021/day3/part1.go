package day3

import (
	"fmt"
	"math/big"
	"strings"
)

const (
	binBase = 2
	square  = 2
)

type reportValue []*big.Int

type diagnosticReport struct {
	reportValue
	reportBitLen int
}

func parseInput(input string) (parsed diagnosticReport, err error) {
	reportedValues := []*big.Int{}
	rawReport := strings.Split(input, "\n")
	reportLength := len(rawReport[0])
	for _, reportedBinary := range rawReport {
		reported := new(big.Int)
		reported, success := reported.SetString(reportedBinary, binBase)
		if !success {
			err = fmt.Errorf("converting value %s failed", reportedBinary)
			break
		}
		reportedValues = append(reportedValues, reported)
	}
	if err == nil {
		parsed = diagnosticReport{reportedValues, reportLength}
	}
	return parsed, err
}

func (d diagnosticReport) getMostCommonBits() *big.Int {
	oneBitCounts := make([]uint, d.reportBitLen)
	for _, reported := range d.reportValue {
		for bitIdx := 0; bitIdx < d.reportBitLen; bitIdx++ {
			oneBitCounts[bitIdx] += reported.Bit(bitIdx)
		}
	}
	reportLen := uint(len(d.reportValue))
	mostCommonBits := new(big.Int)
	for bitIdx, oneBitCount := range oneBitCounts {
		zeroBitCount := reportLen - oneBitCount
		if oneBitCount >= zeroBitCount {
			mostCommonBits.SetBit(mostCommonBits, bitIdx, 1)
		} else if oneBitCount < zeroBitCount {
			mostCommonBits.SetBit(mostCommonBits, bitIdx, 0)
		}
	}
	return mostCommonBits
}

// So, epsilon is supposed to equal ~gamma. I forgot how finicky Go can be with
// binary numbers and signing so this xor hack is required to do the not
// operation that I actually want to do here.
func gammaToEpsilon(gamma *big.Int, bitLen int) *big.Int {
	notOperand := big.NewInt(1)
	for shiftIdx := 0; shiftIdx < bitLen; shiftIdx++ {
		notOperand.Mul(notOperand, big.NewInt(square))
	}
	notOperand.Sub(notOperand, big.NewInt(1))
	epsilon := new(big.Int)
	return epsilon.Xor(gamma, notOperand)
}

// SolvePart1 solves part 1 and returns the power consumption.
func SolvePart1(input string) (int, error) {
	fullReport, err := parseInput(input)
	powerConsumption := new(big.Int)
	var ans int
	if err == nil {
		gamma := fullReport.getMostCommonBits()
		epsilon := gammaToEpsilon(gamma, fullReport.reportBitLen)
		ans = int(powerConsumption.Mul(gamma, epsilon).Int64())
	}
	return ans, err
}
