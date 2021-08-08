package tasks

import (
	"reflect"
	"testing"
)

func TestTribonacci(t *testing.T) {
	var tests = []struct {
		num int
		exp int
	}{
		{
			4, 4,
		},
		{
			25, 1389537,
		},
		{
			34, 334745777,
		},
	}
	for _, tc := range tests {
		res := tribonacci(tc.num)
		if !reflect.DeepEqual(res, tc.exp) {
			t.Errorf("Tribonacci(%d): Expected %d, got %d", tc.num,
				tc.exp, res)
		}
	}

}
