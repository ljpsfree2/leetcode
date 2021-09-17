package tasks

import "fmt"

type P struct {
	row int
	col int
}

func isValidSudoku(board [][]byte) bool {
	var nums map[byte][]P = make(map[byte][]P)
	var num byte
	var is_invalid_sudoku bool

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != []byte(".")[0] {
				p := P{
					row: i,
					col: j,
				}
				num = board[i][j]
				if val, ok := nums[num]; ok {
					nums[num] = append(val, p)
				} else {
					nums[num] = []P{p}
				}
			}
		}

	}

	for key, p_list := range nums {
		// check rows
		for i := 0; i < len(p_list); i++ {
			p1 := p_list[i]
			for j := i + 1; j < len(p_list); j++ {
				p2 := p_list[j]
				if p1.row == p2.row || p1.col == p2.col {
					// two point in the same row or col
					fmt.Println("same row or col: ", key, p1, p2)
					is_invalid_sudoku = true
				}
				if p1.row/3 == p2.row/3 && p1.col/3 == p2.col/3 {
					// two point in the same box
					fmt.Println("same box: ", p1.row/3, p2.row/3, p1.col/3, p2.col/3)
					fmt.Println("same box: ", key, p1, p2)
					is_invalid_sudoku = true
				}
			}
			if is_invalid_sudoku {
				return false
			}
		}
	}
	return true
}
