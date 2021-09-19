package tasks

import "fmt"

type Path struct {
	avg_num int
	l       int
	r       int
	length  int
}

func TrafficJam(n int, k int, ai []int) []int {

	var path_avg [][]int = make([][]int, n)

	for i := 0; i < n; i++ {
		path_avg[i] = make([]int, n)
	}

	var max_num, avg_num int
	var result, path Path
	for i := 0; i < n; i++ {

		path_avg[i][i] = ai[i]

		for j := i + 1; j < n; j++ {

			avg_num = (path_avg[i][j-1] + ai[j]) / 2
			path_avg[i][j] = avg_num
			path = Path{
				avg_num: avg_num,
				l:       i,
				r:       j,
			}

			if avg_num == max_num {
				if path.r-path.l+1 > result.r-result.l+1 {
					result = path
				}
			} else if avg_num > max_num {
				result = path
				max_num = avg_num
			}
		}
	}

	fmt.Println(ai, "ad\n\nad")
	for i := 1; i < n; i++ {
		fmt.Println(path_avg[i])
	}

	return []int{result.l, result.r}

}
