package tasks

import "fmt"

/**
 * Definition for a binary tree node.
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var m map[int]map[int][]int

func insert(arr []int, position int, val int) []int {
	var result []int

	if position >= len(arr) {
		// insert element at the end
		result = append(arr, val)
	} else {
		rear := append([]int{}, arr[position:]...)
		result = append(arr[0:position], val)
		result = append(result, rear...)
	}

	return result
}

func insertOrderly(arr []int, val int) []int {
	i := 0
	for i < len(arr) && val > arr[i] {
		i = i + 1
	}
	arr = insert(arr, i, val)
	return arr
}

func tree_dfs(node *TreeNode, row int, col int) {
	if node == nil {
		return
	}

	if m[col] == nil {
		m[col] = make(map[int][]int)
	}
	if m[col][row] == nil {
		m[col][row] = []int{node.Val}
	} else {
		// insert new value to right position, A->Z
		m[col][row] = insertOrderly(m[col][row], node.Val)
	}

	tree_dfs(node.Left, row+1, col-1)
	tree_dfs(node.Right, row+1, col+1)
}

func verticalTraversal(root *TreeNode) [][]int {
	m = make(map[int]map[int][]int)
	tree_dfs(root, 0, 0)

	// convert map to arrays
	// 1. sort col index
	// 2. move row data to array by sorted col index.
	var cols []map[int][]int // the array to store row data
	var col_indexes []int
	fmt.Println(m)

	// prepare sorted col indexes
	for col_idx := range m {
		// find correct position in cols
		col_indexes = insertOrderly(col_indexes, col_idx)
	}
	fmt.Println(col_indexes)

	// move row data to array by sorted col index
	for i := 0; i < len(col_indexes); i++ {
		cols = append(cols, m[col_indexes[i]])
	}

	var result = make([][]int, len(cols))
	var row_map map[int][]int
	fmt.Println(cols)
	// merge rows
	for i := 0; i < len(cols); i++ {
		/*
			data structure
			row_map -> {row_1: [1,2], row_3:[3,4]}
		*/

		row_map = cols[i]
		var row_arr []int
		var row_indexes []int
		fmt.Println(row_map)

		// prepare sorted row indexes
		for row_idx := range row_map {
			// find correct position in cols
			row_indexes = insertOrderly(row_indexes, row_idx)
		}

		// merge row data by ordered row indexes
		for i := 0; i < len(row_indexes); i++ {
			row_arr = append(row_arr, row_map[row_indexes[i]]...)
		}

		result[i] = row_arr
	}

	return result
}
