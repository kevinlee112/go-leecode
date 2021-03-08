package __100

/**
实现 strStr() 函数。

给定一个 haystack 字符串和一个 needle 字符串，在 haystack 字符串中找出 needle 字符串出现的第一个位置 (从0开始)。如果不存在，则返回  -1。

示例 1:

输入: haystack = "hello", needle = "ll"
输出: 2
示例 2:

输入: haystack = "aaaaa", needle = "bba"
输出: -1
 */

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	if len(needle) > len(haystack) {
		return -1
	}
	i, j, index := 0, 0, -1
	for i<len(haystack) && j < len(needle) {
		if haystack[i] == needle[j] {
			if index == -1 {
				index = i
			}
			if j == len(needle) -1 {
				return index
			}
			j++
		}else {
			if index != -1 {
				i = index
				index = -1
				j=0
			}
		}
		i++
	}
	return -1
}
