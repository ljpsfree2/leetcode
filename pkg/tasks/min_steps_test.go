package tasks

import (
	"reflect"
	"testing"
)

func TestMinSteps(t *testing.T) {
	var tests = []struct {
		n   int
		exp int
	}{
		{
			1,
			0,
		},
		{
			2,
			2,
		},
		{
			3,
			3,
		},
		{
			4,
			4,
		},
		{
			5,
			5,
		},
		{
			9,
			6,
		},
	}
	for _, tc := range tests {
		res := minSteps(tc.n)
		if !reflect.DeepEqual(res, tc.exp) {
			t.Errorf("minSteps(%d): Expected %d, got %d", tc.n,
				tc.exp, res)
		}
	}

}
