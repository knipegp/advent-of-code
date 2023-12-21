package day1

import (
	"testing"
)

func TestGetCalibrationTotal(t *testing.T) {
	calibrationValues := []string{
		"1abc2",
		"pqr3stu8vwx",
		"a1b2c3d4e5f",
		"treb7uchet",
	}
	expectedTotal := 142
	var calibrationTotal int
	var err error
	if calibrationTotal, err = GetCalibrationTotalDigits(calibrationValues); err != nil {
		panic(err)
	}
	if calibrationTotal != expectedTotal {
		t.Errorf("Reported calibration value is %d, expecting %d", calibrationTotal, expectedTotal)
	}
}

func TestGetCalibrationPart2(t *testing.T) {
	calibrationValues := []string{
		"two1nine",
		"eightwothree",
		"abcone2threexyz",
		"xtwone3four",
		"4nineeightseven2",
		"zoneight234",
		"7pqrstsixteen",
	}
	expectedTotal := 281
	var calibrationTotal int
	var err error
	if calibrationTotal, err = GetCalibrationTotalTokens(calibrationValues); err != nil {
		panic(err)
	}
	if calibrationTotal != expectedTotal {
		t.Errorf("Reported calibration value is %d, expecting %d", calibrationTotal, expectedTotal)
	}
}
