package tasks

import (
	"reflect"
	"testing"
)

func TestKWeakestRows(t *testing.T) {
	var tests = []struct {
		mat [][]int
		k   int
		exp []int
	}{
		{
			[][]int{{1, 1, 0, 0, 0},
				{1, 1, 1, 1, 0},
				{1, 0, 0, 0, 0},
				{1, 1, 0, 0, 0},
				{1, 1, 1, 1, 1}},
			3,
			[]int{2, 0, 3}},
		{
			[][]int{{1, 0, 0, 0},
				{1, 1, 1, 1},
				{1, 0, 0, 0},
				{1, 0, 0, 0}},
			2,
			[]int{0, 2}},

		{
			[][]int{
				{1, 1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1, 1}},
			1,
			[]int{0}},
		{
			[][]int{
				{1, 0},
				{1, 0},
				{1, 0},
				{1, 1}},
			4,
			[]int{0, 1, 2, 3}},
	}

	for _, tc := range tests {
		res := kWeakestRows(tc.mat, tc.k)
		if !reflect.DeepEqual(res, tc.exp) {
			t.Errorf("kWeakestRows(%d, %d): Expected %d, got %d", tc.mat,
				tc.k, tc.exp, res)
		}
	}
}
