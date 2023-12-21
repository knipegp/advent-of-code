package day3

import (
	"fmt"
	"strconv"
)

type Schematic [][]byte

type Coordinate struct {
	x int
	y int
}

type Edge struct {
	a Coordinate
	b Coordinate
}

func (e Edge) GetMiddleCoordinates() []Coordinate {
	middleCoordinates := []Coordinate{}
	xDiff := e.a.x - e.b.x
	if xDiff > 0 {
		for index := e.b.x; index < e.a.x; index++ {
			middleCoordinates = append(middleCoordinates, Coordinate{index, e.b.y})
		}
	} else if xDiff < 0 {
		for index := e.a.x; index < e.b.x; index++ {
			middleCoordinates = append(middleCoordinates, Coordinate{index, e.b.y})
		}
	}
	yDiff := e.a.y - e.b.y
	if yDiff > 0 {
		for index := e.b.y; index < e.a.y; index++ {
			middleCoordinates = append(middleCoordinates, Coordinate{index, e.b.x})
		}
	} else if yDiff < 0 {
		for index := e.a.y; index < e.b.y; index++ {
			middleCoordinates = append(middleCoordinates, Coordinate{index, e.b.x})
		}
	}
	return middleCoordinates
}

func (e Edge) MiddleContainsSymbol(schematic Schematic) bool {
	for _, coordinate := range e.GetMiddleCoordinates() {
		schematicChar := schematic[coordinate.y][coordinate.x]
		if schematicChar != '.' {
			if _, err := strconv.Atoi(string(schematicChar)); err != nil {
				return true
			}
		}
	}
	return false
}

type SchematicNumber struct {
	number     int
	edges      [4]Edge
	partNumber bool
}

func NewSchematicNumber() {

}
