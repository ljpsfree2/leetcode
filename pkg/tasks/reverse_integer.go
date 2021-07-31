package tasks

import (
	"math"
)

func reverse(x int) int {
	rev := 0
	num := 0
	flag := 1

	if x < 0 {
		flag = -1
	}

	x = int(math.Abs(float64(x)))

	for x > 0 {
		num = x % 10
		rev = rev*10 + num
		x = x / 10
	}

	rev = rev * flag

	if rev > int(math.Pow(2, 31)) || rev < int(-1*math.Pow(2, 31)) {
		rev = 0
	}

	return rev
}

func reverse2(x int) int {
	var arr [32]int
	flag := false
	total := 0

	if x < 0 {
		flag = true
	}

	num := int(math.Abs(float64(x)))

	for num > 0 {
		arr[total] = num % 10
		total = total + 1
		num = num / 10
	}

	/*
		tmp := 0
			for i := 0; i < total/2; i++ {
				tmp = arr[i]
				arr[i] = arr[total-(i+1)]
				arr[total-(i+1)] = tmp
				fmt.Println(tmp, arr[i], total-(i+1))
			}
			fmt.Println(arr)
	*/

	result := 0
	for i := 0; i < total; i++ {
		result = result + arr[i]*int(math.Pow(10, float64(total-(i+1))))
	}

	if flag {
		result = result * -1
	}

	if result > int(math.Pow(2, 31)) || result < int(-1*math.Pow(2, 31)) {
		result = 0
	}

	return result
}

func Reverse(x int) int {
	return reverse(x)
}
