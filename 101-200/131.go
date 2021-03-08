package _01_200

import "fmt"

/**
给定一个字符串 s，将 s 分割成一些子串，使每个子串都是回文串。

返回 s 所有可能的分割方案。

示例:

输入: "aab"
输出:
[
  ["aa","b"],
  ["a","a","b"]
]
 */
func Partition(s string) [][]string {
	var res [][]string
	backtrace(s, 0, []string{}, &res)
	return res
}
func backtrace(s string, start int, tmp []string, res *[][]string)  {
	if start == len(s) {
		t := make([]string, len(tmp)) // 新建一个和temp等长的切片
		copy(t, tmp)                  // temp还要在递归中继续被修改，不能将它的引用推入res
		*res = append(*res, t)
		return
	}
	for i:=start;i<len(s);i++ {
		if isPartition(s[start:i+1]) {
			tmp = append(tmp, s[start:i+1])
			fmt.Println(tmp)
			backtrace(s, i+1, tmp, res)
			tmp = tmp[:len(tmp)-1]
		}
	}
}

func isPartition(s string) bool {
	l, r := 0, len(s)-1
	for l < r {
		if s[l] != s[r] {
			return false
		}
		r--
		l++
	}
	return true
}







