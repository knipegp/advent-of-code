package day3

import (
	"fmt"
	"math/big"
	"strings"
)

type reportValue []*big.Int

type diagnosticReport struct {
	reportValue
	reportBitLen int
}

func parseInput(input string) diagnosticReport {
	reportedValues := []*big.Int{}
	rawReport := strings.Split(input, "\n")
	reportLength := len(rawReport[0])
	for _, reportedBinary := range rawReport {
		reported := new(big.Int)
		reported, success := reported.SetString(reportedBinary, 2)
		if !success {
			panic(fmt.Errorf("Converting value %s failed", reportedBinary))
		}
		reportedValues = append(reportedValues, reported)
	}
	return diagnosticReport{reportedValues, reportLength}
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
		if oneBitCount > reportLen/2 {
			mostCommonBits.SetBit(mostCommonBits, bitIdx, 1)
		} else if oneBitCount < reportLen/2 {
			mostCommonBits.SetBit(mostCommonBits, bitIdx, 0)
		} else {
			panic(fmt.Errorf("1 and 0 bit counts are equal for bit %d", bitIdx))
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
		notOperand.Mul(notOperand, big.NewInt(2))
	}
	notOperand.Sub(notOperand, big.NewInt(1))
	epsilon := new(big.Int)
	return epsilon.Xor(gamma, notOperand)
}

// SolvePart1 solves part 1 and returns the power consumption.
func SolvePart1(input string) int {
	fullReport := parseInput(input)
	gamma := fullReport.getMostCommonBits()
	epsilon := gammaToEpsilon(gamma, fullReport.reportBitLen)
	powerConsumption := new(big.Int)
	return int(powerConsumption.Mul(gamma, epsilon).Int64())
}
