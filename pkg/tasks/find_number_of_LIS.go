package tasks

/*func findNumberOfLIS(nums []int) int {
	var num, max_num, max_length, j int

	for i := 0; i < len(nums); i++ {
		for j = 0; nums[j] > nums[i]; i++ {
		}
		if j > max_length {
			max_length = j
		}
	}
}
*/
func findNumberOfLIS(nums []int) (ans int) {
	maxLen := 0
	n := len(nums)
	dp := make([]int, n)
	cnt := make([]int, n)
	for i, x := range nums {
		dp[i] = 1
		cnt[i] = 1
		for j, y := range nums[:i] {
			if x > y {
				if dp[j]+1 > dp[i] {
					dp[i] = dp[j] + 1
					cnt[i] = cnt[j] // 重置计数
				} else if dp[j]+1 == dp[i] {
					cnt[i] += cnt[j]
				}
			}
		}
		if dp[i] > maxLen {
			maxLen = dp[i]
			ans = cnt[i] // 重置计数
		} else if dp[i] == maxLen {
			ans += cnt[i]
		}
	}
	return
}
