package tasks

import (
	"reflect"
	"testing"
)

func TestTrafficJam(t *testing.T) {
	var tests = []struct {
		n   int
		k   int
		ai  []int
		exp []int
	}{
		{
			9,
			3,
			[]int{1, 3, 2, 1, 1, 2, 2, 2, 2},
			[]int{5, 8},
		},
	}

	for _, tc := range tests {
		res := TrafficJam(tc.n, tc.k, tc.ai)
		if !reflect.DeepEqual(res, tc.exp) {
			t.Errorf("TrafficJam(%d, %d, %d): Expected %d, got %d", tc.n, tc.k,
				tc.ai, tc.exp, res)
		}
	}

}
