package _01_400

/**
编写一个函数，以字符串作为输入，反转该字符串中的元音字母。

 

示例 1：

输入："hello"
输出："holle"
示例 2：

输入："leetcode"
输出："leotcede"
 */

func reverseVowels(s string) string {
	b := []byte(s)
	for l, r := 0, len(s)-1;l<r; {
		for l < r && !isVowels(b[l]){
			l++
		}
		for l<r && !isVowels(b[r]) {
			r--
		}
		b[l], b[r] = b[r], b[l]
		l++
		r--
	}
	return string(b)
}

func isVowels(ch byte) bool {
	switch ch {
	case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
		return true
	default:
		return false
	}
}
