package hard

func findSubstring(s string, words []string) []int {
	var result []int // 最终结果
	// 特判
	if len(s) == 0 || len(words) == 0 || len(words[0]) *len(words) > len(s) {
		return result
	}
	wordLen, wordNum, ls := len(words[0]), len(words), len(s)
	totalLen := wordLen * wordNum // 子串总长度
	if totalLen > ls {
		return result
	}

	// 统计words中每个单词出现的次数
	need := make(map[string]int)
	for _, word := range words {
		need[word]++
	}

	for start := 0; start < wordLen; start++ {
		// 窗口左边，窗口右边，词频匹配的单词数
		left, right, match := start, start, 0
		// 窗口中每个单词出现的次数
		window := make(map[string]int)
		for right+wordLen <= ls {
			rightWord := s[right : right+wordLen] // 当前新加入的单词
			right += wordLen
			if need[rightWord] > 0 {
				window[rightWord]++
				if window[rightWord] == need[rightWord] {
					match++
				}
			}
			// 如果满足了长度，判断是否满足词频
			if right-left == totalLen {
				// 词频也匹配，加入结果
				if match == len(need) {
					result = append(result, left)
				}
				leftWord := s[left : left+wordLen] // 当前新移出的单词
				left += wordLen
				if need[leftWord] > 0 {
					if window[leftWord] == need[leftWord] {
						match--
					}
					window[leftWord]--
				}
			}
		}
	}
	return result
}