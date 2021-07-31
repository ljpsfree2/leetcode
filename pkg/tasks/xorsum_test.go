package tasks

import (
	"testing"
)

func TestSubsetXorSum(t *testing.T) {
	nums := []int{5, 1, 6}
	result := SubsetXORSum(nums)
	if result != 28 {
		t.Fatalf("num: 123, expected 321, result: %d", result)
	}
}
