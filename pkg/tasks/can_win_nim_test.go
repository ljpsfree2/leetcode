package tasks

import (
	"reflect"
	"testing"
)

func TestCanWinNim(t *testing.T) {
	var tests = []struct {
		n   int
		exp bool
	}{
		{
			4,
			false,
		},
		{
			5,
			true,
		},
	}
	for _, tc := range tests {
		res := canWinNim(tc.n)
		if !reflect.DeepEqual(res, tc.exp) {
			t.Errorf("FindIntegers(%d): Expected %t, got %t", tc.n,
				tc.exp, res)
		}
	}

}
