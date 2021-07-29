package tasks

func TwoSum(nums []int, target int) []int {
	var result = []int{0, 0}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				result[0] = i
				result[1] = j
				return result
			}
		}
	}
	return result

}

func TwoSum2(nums []int, target int) []int {
	// hashTable 记录所有已经遍历过的数，以余数为key，index为value
	// 算法复杂度O（N），多了hashTable的存储
	hashTable := map[int]int{}
	for i, x := range nums {
		if p, ok := hashTable[target-x]; ok {
			return []int{p, i}
		}
		hashTable[x] = i
	}
	return nil
}
