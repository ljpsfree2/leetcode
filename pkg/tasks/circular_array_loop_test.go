package tasks

import (
	"reflect"
	"testing"
)

func TestCircularArrayLoop(t *testing.T) {
	var tests = []struct {
		nums []int
		exp  bool
	}{
		{
			[]int{2, -1, 1, 2, 2},
			true,
		},
		{
			[]int{-1, 2},
			false,
		},
		{
			[]int{-1, -2, -3, -4, -5},
			false,
		},
		{
			[]int{-2, 1, -1, -2, -2},
			false,
		},
		{
			[]int{3, 1, 2},
			true,
		},
	}
	for _, tc := range tests {
		res := circularArrayLoop(tc.nums)
		if !reflect.DeepEqual(res, tc.exp) {
			t.Errorf("circularArrayLoop(%d): Expected %t, got %t", tc.nums,
				tc.exp, res)
		}
	}

}
