package tasks

func kWeakestRows(mat [][]int, k int) []int {
	var result []int
	if mat == nil {
		return result
	}
	var row_idxes = make([]int, len(mat))

	for i := range row_idxes {
		row_idxes[i] = i
	}

	for col := 0; col < len(mat[0]); col++ {
		var row_idxes_left []int
		for _, row_idx := range row_idxes {
			if mat[row_idx][col] == 0 {
				result = append(result, row_idx)
				if len(result) >= k {
					return result
				}
			} else {
				row_idxes_left = append(row_idxes_left, row_idx)
			}
		}
		row_idxes = row_idxes_left
	}

	for len(result) < k {
		result = append(result, row_idxes[0:k-len(result)]...)
	}
	return result
}
