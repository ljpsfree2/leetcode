package tasks

import (
	"testing"
)

func TestTwoSum(t *testing.T) {
	nums := []int{2, 3, 4}
	target := 6
	var result []int = TwoSum(nums, target)

	if result[0] != 0 || result[1] != 2 {
		t.Fatalf("nums: [2,3,4], expected [0,2]")
	}

}
