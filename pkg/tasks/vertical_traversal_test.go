package tasks

import (
	"reflect"
	"testing"
)

func make_tree(node_nums []int) TreeNode {
	var queue []*TreeNode
	var top, node *TreeNode
	i := 0
	qi := 0
	root := TreeNode{
		Val:   node_nums[0],
		Left:  nil,
		Right: nil,
	}
	queue = append(queue, &root)
	for i < len(node_nums) {
		top = queue[qi]
		i = i + 1
		if i < len(node_nums) && node_nums[i] >= 0 {
			node = &TreeNode{
				Val:   node_nums[i],
				Left:  nil,
				Right: nil,
			}
			top.Left = node
			queue = append(queue, node)
		}
		i = i + 1
		if i < len(node_nums) && node_nums[i] > 0 {
			node = &TreeNode{
				Val:   node_nums[i],
				Left:  nil,
				Right: nil,
			}
			top.Right = node
			queue = append(queue, node)
		}
		qi = qi + 1
	}
	return root
}

func tree_to_arr(root *TreeNode) []int {
	var cur_arr []*TreeNode
	var next_arr []*TreeNode
	var res []int

	cur_arr = []*TreeNode{root}
	has_next_level := true
	for has_next_level {
		has_next_level = false
		next_arr = []*TreeNode{}
		for _, node_p := range cur_arr {
			var val int
			var left, right *TreeNode
			if node_p == nil {
				val = -1
				left = nil
				right = nil
			} else {
				val = (*node_p).Val
				left = (*node_p).Left
				right = (*node_p).Right
			}
			res = append(res, val)
			if left != nil || right != nil {
				has_next_level = true
			}
			next_arr = append(next_arr, left, right)
		}

		cur_arr = next_arr
	}
	return res
}

func TestTreeMaker(t *testing.T) {
	var tests = []struct {
		nums []int
		exp  []int
	}{
		{
			[]int{3, 9, 20, -1, -1, 15, 7},
			[]int{3, 9, 20, -1, -1, 15, 7},
		},
		{
			[]int{1, 2, 3, 4, 5, 6, 7},
			[]int{1, 2, 3, 4, 5, 6, 7},
		},
		{
			[]int{1, 2, 3, 4, 6, 5, 7},
			[]int{1, 2, 3, 4, 6, 5, 7},
		},
	}
	for _, tc := range tests {
		tree := make_tree(tc.nums)
		arr := tree_to_arr(&tree)
		if !reflect.DeepEqual(arr, tc.exp) {
			t.Errorf("make_tree(%d): Expected %d, got %d", tc.nums, tc.exp, arr)
		}
	}

}

func TestVerticalTraversal(t *testing.T) {
	var tests = []struct {
		input []int
		exp   [][]int
	}{
		{
			[]int{3, 9, 20, -1, -1, 15, 7},
			[][]int{{9}, {3, 15}, {20}, {7}},
		},
		{
			[]int{1, 2, 3, 4, 5, 6, 7},
			[][]int{{4}, {2}, {1, 5, 6}, {3}, {7}},
		},
		{
			[]int{1, 2, 3, 4, 6, 5, 7},
			[][]int{{4}, {2}, {1, 5, 6}, {3}, {7}},
		},
		{
			[]int{0, 1, -1, -1, 2, 6, 3, -1, -1, -1, 4, -1, 5},
			[][]int{{1, 6}, {0, 2}, {3}, {4}, {5}},
		},
		{
			[]int{3, 1, 4, 0, 2, 2},
			[][]int{{0}, {1}, {3, 2, 2}, {4}},
		},
	}
	for _, tc := range tests {
		tree := make_tree(tc.input)
		res := verticalTraversal(&tree)
		if !reflect.DeepEqual(res, tc.exp) {
			t.Errorf("verticalTraversal(%d): Expected %d, got %d", tc.input,
				tc.exp, res)
		}
	}
}

func TestInsert(t *testing.T) {
	var tests = []struct {
		arr      []int
		position int
		val      int
		exp      []int
	}{
		{
			[]int{1, 2, 3},
			0,
			100,
			[]int{100, 1, 2, 3},
		},
		{
			[]int{1, 2, 3},
			3,
			100,
			[]int{1, 2, 3, 100},
		},
		{
			[]int{1, 2, 3},
			1,
			100,
			[]int{1, 100, 2, 3},
		},
	}
	for _, tc := range tests {
		res := insert(tc.arr, tc.position, tc.val)
		if !reflect.DeepEqual(res, tc.exp) {
			t.Errorf("insert(%d, %d, %d): Expected %d, got %d", tc.arr,
				tc.position, tc.val, tc.exp, res)
		}
	}
}

func TestInsertOrderly(t *testing.T) {
	var tests = []struct {
		arr []int
		val int
		exp []int
	}{
		{
			[]int{1, 2, 3},
			100,
			[]int{1, 2, 3, 100},
		},
		{
			[]int{1, 5, 3},
			6,
			[]int{1, 5, 3, 6},
		},
		{
			[]int{1, 29, 320},
			100,
			[]int{1, 29, 100, 320},
		},
	}
	for _, tc := range tests {
		res := insertOrderly(tc.arr, tc.val)
		if !reflect.DeepEqual(res, tc.exp) {
			t.Errorf("insert(%d, %d): Expected %d, got %d", tc.arr,
				tc.val, tc.exp, res)
		}
	}
}
