package tasks

func minDistance(word1 string, word2 string) int {
	var i, j int = 0
	var small_word, large_word string
	var path [500]string

	small_word = word1
	small_size := len(small_word)
	large_word = word2
	large_size := len(large_word)

	if small_size > large_size {
		large_size = small_size + large_size
		small_size = large_size - small_size
		large_size = large_size - small_size
		large_word = word2
		small_word = word1
	}

	path[0] = small_word[0]
	for i >= 0 {
		j = i
		for j < large_size && small_word[i] != large_word[j] {
			j = j + 1
		}
		if j < large_size {

		}
	}
	len1 = len(word1)
	len2 = len(word2)

}
