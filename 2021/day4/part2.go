package day4

import "fmt"

func removeBoards(boards []bingoBoard, idxs []int) []bingoBoard {
	for removeIdx, boardIdx := range idxs {
		boards = append(boards[:boardIdx-removeIdx], boards[boardIdx-removeIdx+1:]...)
	}
	return boards
}

// SolvePart2 calculates the final bingo score for the last completed board.
func SolvePart2(input string) (score int, err error) {
	var boards []bingoBoard
	var moves []int
	moves, boards, err = parseInput(input)
	if err == nil {
		var validLastWinner bool
		for _, move := range moves {
			removeIdxs := []int{}
			for boardIdx := range boards {
				boards[boardIdx].playMove(move)
				if boards[boardIdx].isComplete() {
					removeIdxs = append(removeIdxs, boardIdx)
				}
			}
			if len(boards) > 1 && len(removeIdxs) > 0 {
				boards = removeBoards(boards, removeIdxs)
			} else if len(boards) == 1 && len(removeIdxs) == 1 {
				score, err = boards[0].getScore(move)
				validLastWinner = true
				break
			}
		}
		if !validLastWinner {
			err = fmt.Errorf("Could not complete the final winning board")
		}
	}
	return score, err
}
