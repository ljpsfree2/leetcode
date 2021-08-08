package tasks

func tribonacci(n int) int {

	var num1, num2, num3, num int
	num1 = 0
	num2 = 1
	num3 = 1

	switch n {
	case 0:
		return 0
	case 1:
		return 1
	case 2:
		return 1
	}

	for i := 3; i <= n; i++ {
		num = num1 + num2 + num3
		num1 = num2
		num2 = num3
		num3 = num
	}

	return num
}
