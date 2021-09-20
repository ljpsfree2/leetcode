package tasks

import (
	"reflect"
	"testing"
)

func TestFindNumberOfLIS(t *testing.T) {
	var tests = []struct {
		n   []int
		exp int
	}{
		{
			[]int{1, 3, 5, 4, 7},
			2,
		},
		{
			[]int{2, 2, 2, 2, 2},
			5,
		},
	}
	for _, tc := range tests {
		res := findNumberOfLIS(tc.n)
		if !reflect.DeepEqual(res, tc.exp) {
			t.Errorf("findNumberOfLIS(%d): Expected %d, got %d", tc.n,
				tc.exp, res)
		}
	}

}
