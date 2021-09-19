package tasks

func minSteps(n int) int {
	var steps, max_copy_num int

	steps = 0

	if n == 1 {
		return 0
	}

	for i := n / 2; i >= 1; i-- {
		if n%i == 0 {
			max_copy_num = i
			break
		}
	}
	steps = minSteps(max_copy_num) + n/max_copy_num
	return steps
}
