package __100

/**

给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。

示例 1:

输入: "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
示例 2:

输入: "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
示例 3:

输入: "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
*/

func lengthOfLongestSubstring(s string) int {
	if len(s) <= 1 {
		return len(s)
	}
	res := 0
	m := make(map[byte]int)
	i, j := 0, 0
	for j < len(s) {
		if v, exit := m[s[j]]; exit {
			if i < v+1 {
				i = v + 1
			}
		}
		m[s[j]] = j

		if res < j-i+1 {
			res = j - i + 1
		}
		j++
	}
	return res
}
