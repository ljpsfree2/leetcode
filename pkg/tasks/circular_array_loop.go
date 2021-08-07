package tasks

import "fmt"

func circularArray(nums []int, start int) bool {
	var loop []int = make([]int, len(nums))
	var orig_nums []int = make([]int, len(nums))
	var num, num_idx, i int

	copy(orig_nums, nums)
	fmt.Println(nums)
	num_idx = start

	num = nums[num_idx]
	i = 0
	for num != 0 {
		// store the sequence number
		loop[i] = num_idx

		// cleanup current position, set num[num_idx] = 0
		// because no 0 in nums
		nums[num_idx] = 0

		// get next position
		i = i + 1
		num_idx = (num_idx + num) % len(nums)
		if num_idx < 0 {
			num_idx = num_idx + len(nums)
		}
		fmt.Println("new", num_idx)
		num = nums[num_idx]
	}

	loop_end := i
	loop_start := 0
	for loop[loop_start] != num_idx {
		loop_start = loop_start + 1
	}

	// verify loop
	fmt.Println(loop_start, loop_end, loop, loop[loop_start:loop_end], len(loop[loop_start:loop_end]))
	if len(loop[loop_start:loop_end]) > 1 {
		is_positive := false
		if orig_nums[loop[loop_start]] > 0 {
			is_positive = true
		}

		for i := loop_start; i < loop_end; i++ {
			fmt.Println("loop", orig_nums, i, loop[i], orig_nums[loop[i]])
			num = orig_nums[loop[i]]
			if is_positive && num < 0 {
				return false
			}
			if !is_positive && num > 0 {
				return false
			}
		}
		return true
	}
	return false
}
func circularArrayLoop(nums []int) bool {
	for i := 0; i < len(nums); i++ {
		// start from everywhere
		var temp_nums []int = make([]int, len(nums))
		copy(temp_nums, nums)
		if circularArray(temp_nums, i) {
			return true
		}
	}
	return false
}
