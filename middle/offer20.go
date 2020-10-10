package middle

import "strings"

func isNumber(s string) bool {
	// 去除首尾空格
	s = strings.TrimSpace(s)

	for i := 0; i < len(s); i++ {
		// 存在 e 或 E, 判断是否为科学计数法
		if s[i] == 'e' || s[i] == 'E' {
			return isSciNum(s[:i],s[i+1:])
		}
	}

	// 否则判断是否为整数或小数
	return isInt(s) || isDec(s)
}

// 是否为科学计数法
func isSciNum(num1,num2 string) bool {
	// e 前后字符串长度为0 是错误的
	if len(num1) == 0 || len(num2) == 0 {
		return false
	}
	// e 后面必须是整数，前面可以是整数或小数
	return (isInt(num1) || isDec(num1)) && isInt(num2)
}

// 判断是否为小数
func isDec(s string) bool {
	if len(s) == 0 {
		return false
	}
	// 去除正负符号
	if s[0] == '+' || s[0] == '-' {
		s = s[1:]
	}
	for i := 0; i < len(s); i++ {
		if s[i] == '.' {
			// 允许小数点前或小数点后没有数字，但不允许同时没有数字
			// 如果小数点后有数字存在，应为无符号整数
			if i == 0 && i != len(s)-1 {
				if s[i+1] == '+' || s[i+1] == '-' {
					return false
				}else{
					return isInt(s[i+1:])
				}
			}
			// 如果小数点前有数字存在，应为整数
			if i != 0 && i == len(s)-1 {
				return isInt(s[:i])
			}

			// 小数点前后数字均存在时，均需符合
			if i != 0 && i != len(s)-1 {
				var b bool
				if s[i+1] == '+' || s[i+1] == '-' {
					b = false
				}else{
					b = isInt(s[i+1:])
				}
				return isInt(s[:i]) && b
			}
			return false
		}
	}
	return false
}

// 判断是否为整数
func isInt(s string) bool {
	if len(s) == 0 {
		return false
	}
	// 去除正负符号
	if s[0] == '+' || s[0] == '-' {
		s = s[1:]
	}
	// 除了符号以外必须还有数字
	if len(s) == 0 {
		return false
	}
	// 出现数字以外的字符是错误的
	for i := 0; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return false
		}
	}
	return true
}