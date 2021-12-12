package day4

import (
	"fmt"
	"strconv"
	"strings"
)

type bingoBoard struct {
	cellValues [5][5]int
	cellMarked [5][5]bool
}

func (b bingoBoard) playMove(move int) {
	found := false
	for rowIdx, row := range b.cellValues {
		var colIdx, colElem int
		for colIdx, colElem = range row {
			if move == colElem {
				found = true
				break
			}
		}
		if found {
			b.cellMarked[rowIdx][colIdx] = true
			break
		}
	}
}

func (b bingoBoard) isComplete() (isComplete bool) {
	for rowIdx, row := range b.cellMarked {
		isColComplete := true
		for _, colMarked := range row {
			if colMarked == false {
				isColComplete = false
				break
			}
		}
		if isColComplete {
			isComplete = true
			break
		}
	}
	if !isColComplete {
		for colIdx := 0; colIdx < 5; colIdx++ {
			isRowComplete := true
			for _, colMarked := range b.cellMarked[:][colIdx] {
				if colMarked == false {
					isColComplete = false
					break
				}
			}
			if isColComplete {
				isComplete = true
				break
			}
		}
	}
}

func splitRow(input string) (rowElem []string, err error) {
	rowElem = []string{}
	splitString := strings.Split(input, " ")
	for _, splitElem := range splitString {
		if splitElem != "" {
			rowElem = append(rowElem, splitElem)
		}
	}
	if len(rowElem) != 5 {
		err = fmt.Errorf("Line %s could not be parsed into 5 values", input)
	}
	return rowElem, err
}

func boardFromLines(rawBoard [5]string) (newBoard bingoBoard, err error) {
	for rowIdx, rawRow := range rawBoard {
		var rowElems []string
		rowElems, err = splitRow(rawRow)
		if err == nil {
			for colIdx, rawVal := range rowElems {
				newBoard.cellValues[rowIdx][colIdx], err = strconv.Atoi(rawVal)
				if err != nil {
					err = fmt.Errorf(
						"Encountered %w while parsing board %s",
						err,
						rawBoard,
					)
					break
				}
			}
		}
		if err != nil {
			break
		}
	}
	return newBoard, err
}

func movesFromString(input string) (moves []int, err error) {
	rawMoves := strings.Split(input, ",")
	moves = make([]int, len(rawMoves))
	for moveIdx, rawMove := range rawMoves {
		moves[moveIdx], err = strconv.Atoi(rawMove)
		if err != nil {
			err = fmt.Errorf("Encountered %w while parsing moves %s", err, input)
			break
		}
	}
	return moves, err
}

func getInputLines(input string) (lines []string, err error) {
	inputLines := strings.Split(input, "\n")
	lines = []string{}
	for _, line := range inputLines {
		if strings.Compare(line, "") != 0 {
			lines = append(lines, line)
		}
	}
	if (len(inputLines)-1)%5 != 0 {
		err = fmt.Errorf(
			"Day 4 solution expects 1 moves line and any number of 5 line bingo boards",
		)
	}
	return lines, err
}

func parseInput(input string) (moves []int, boards []bingoBoard, err error) {
	var inputLines []string
	inputLines, err = getInputLines(input)
	if err == nil {
		moves, err = movesFromString(inputLines[0])
	}
	if err == nil {
		boardLines := inputLines[1:]
		boards = make([]bingoBoard, len(boardLines)/5)
		for boardIdx := 0; boardIdx < len(boards); boardIdx++ {
			singleBoardLines := [5]string{}
			startLineIdx := boardIdx * 5
			copy(singleBoardLines[:], boardLines[startLineIdx:startLineIdx+5])
			var parsedBoard bingoBoard
			parsedBoard, err = boardFromLines(singleBoardLines)
			if err != nil {
				break
			}
			boards[boardIdx] = parsedBoard
		}
	}
	return moves, boards, err
}

// SolvePart1 calculates the final bingo score.
func SolvePart1(input string) (score int, err error) {
	var boards []bingoBoard
	moves, boards, err = parseInput(input)
	return score, err
}
