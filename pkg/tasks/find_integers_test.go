package tasks

import (
	"reflect"
	"testing"
)

func TestFindIntegers(t *testing.T) {
	var tests = []struct {
		n   int
		exp int
	}{
		{
			5,
			5,
		},
	}
	for _, tc := range tests {
		res := FindIntegers(tc.n)
		if !reflect.DeepEqual(res, tc.exp) {
			t.Errorf("FindIntegers(%d): Expected %d, got %d", tc.n,
				tc.exp, res)
		}
	}

}
