package __100

import "math"

/**
给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。

注意：如果 s 中存在这样的子串，我们保证它是唯一的答案。

 

示例 1：
输入：s = "ADOBECODEBANC", t = "ABC"
输出："BANC"

示例 2：
输入：s = "a", t = "a"
输出："a"
*/
func minWindow(s string, t string) string {
	need, window := map[byte]int{}, map[byte]int{}
	for i := range t {
		need[t[i]]++
	}
	valid := 0
	left, right := 0, 0
	index, length := 0, math.MaxInt32
	for right < len(s) {
		tempAdd := s[right]
		right++
		if _, ok := need[tempAdd]; ok {
			window[tempAdd]++
			if window[tempAdd] == need[tempAdd] {
				valid++
			}
		}
		for valid == len(need) {
			if right-left < length {
				index = left
				length = right - left
			}

			tempDel := s[left]
			left++

			if _, ok := need[tempDel]; ok {
				window[tempDel]--
				if window[tempDel] < need[tempDel] {
					valid--
				}
			}
		}
	}

	if length == math.MaxInt32 {
		return ""
	}
	return s[index : index+length]
}
