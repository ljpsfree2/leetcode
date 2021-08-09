package tasks

import (
	"reflect"
	"testing"
)

func TestCircularArrayLoop(t *testing.T) {
	var tests = []struct {
		n      int
		primes []int
		exp    int
	}{
		{
			12,
			[]int{2, 7, 13, 19},
			32,
		},
		{
			1,
			[]int{2, 3, 5},
			1,
		},
	}
	for _, tc := range tests {
		res := nthSuperUglyNumber(tc.n, tc.primes)
		if !reflect.DeepEqual(res, tc.exp) {
			t.Errorf("circularArrayLoop(%d, %d): Expected %d, got %d", tc.n,
				tc.primes, tc.exp, res)
		}
	}

}
