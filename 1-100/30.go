package __100

/**
给定一个字符串 s 和一些长度相同的单词 words。找出 s 中恰好可以由 words 中所有单词串联形成的子串的起始位置。

注意子串要与 words 中的单词完全匹配，中间不能有其他字符，但不需要考虑 words 中单词串联的顺序。

示例 1：

输入：
  s = "barfoothefoobarman",
  words = ["foo","bar"]
输出：[0,9]
解释：
从索引 0 和 9 开始的子串分别是 "barfoo" 和 "foobar" 。
输出的顺序不重要, [9,0] 也是有效答案。
示例 2：

输入：
  s = "wordgoodgoodgoodbestword",
  words = ["word","good","best","word"]
输出：[]
 */
func findSubstring(s string, words []string) []int {
	var res []int
	if len(s)==0 || len(words)==0 || len(words)*len(words[0]) > len(s) {
		return res
	}
	wordLen, wordNum, slen := len(words[0]), len(words), len(s)
	totalLen := wordLen * wordNum

	need := make(map[string]int)
	for _,word:=range words{
		need[word]++
	}
	for i:=0;i<wordLen;i++ {
		left, right, match := i, i, 0
		window := make(map[string]int)
		for right + wordLen <= slen {
			rightWord := s[right : right+wordLen]
			right += wordLen
			if need[rightWord] > 0 {
				window[rightWord]++
				if window[rightWord] == need[rightWord] {
					match++
				}
			}
			if right-left == totalLen {
				if match == len(need) {
					res = append(res, left)
				}
				leftWord := s[left : left+wordLen]
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
	return res
}