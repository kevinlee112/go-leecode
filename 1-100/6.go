package __100

import "strings"

/**
将一个给定字符串 s 根据给定的行数 numRows ，以从上往下、从左到右进行 Z 字形排列。
比如输入字符串为 "PAYPALISHIRING" 行数为 3 时，排列如下：

P   A   H   N
A P L S I I G
Y   I   R
之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："PAHNAPLSIIGYIR"。

 */
func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	res := make([]string, numRows)
	period := numRows * 2 -2
	for i:=0;i<len(s);i++ {
		mod := i % period
		if mod < numRows {
			res[mod] += string(s[i])
		}else {
			res[period - mod] += string(s[i])
		}
	}
	return strings.Join(res, "")
}