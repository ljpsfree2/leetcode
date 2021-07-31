package tasks

var result int
var nums []int

func dfs(val int, idx int) {
	if idx == len(nums) {
		// 退出递归
		result = result + val
		return
	}
	dfs(val^nums[idx], idx+1)
	dfs(val, idx+1)
}

func SubsetXORSum(arr []int) int {
	nums = arr
	dfs(0, 0)
	return result
}
