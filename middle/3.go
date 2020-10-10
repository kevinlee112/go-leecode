package middle

/**
无重复字符的最长子串
 */

func LengthOfLongestSubstring(s string) int {
	result := 0
	if len(s) <= 1 {
		return result
	}
	m := make(map[byte]int)
	for i, j:=0,0;j<len(s);j++ {
		if v, exit := m[s[j]]; exit {
			if i < v + 1 {
				i = v + 1
			}
		}
		m[s[j]] = j
		if result < j - i + 1 {
			result = j - i + 1
		}
	}
	return result
}

