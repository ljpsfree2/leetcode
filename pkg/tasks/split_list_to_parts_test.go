package tasks

import (
	"reflect"
	"testing"
)

func TestSplitListToParts(t *testing.T) {
	var tests = []struct {
		n   int
		exp int
	}{}
	for _, tc := range tests {
		res := splitListToParts(tc.n)
		if !reflect.DeepEqual(res, tc.exp) {
			t.Errorf("FindIntegers(%d): Expected %d, got %d", tc.n,
				tc.exp, res)
		}
	}

}
