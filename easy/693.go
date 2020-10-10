package easy

import "strconv"

func hasAlternatingBits(n int) bool {
	binary := strconv.FormatInt(int64(n), 2)
	for i := 0; i < len(binary) - 1; i++ {
		if binary[i] == binary[i + 1] {
			return false
		}
	}
	return true
}

func hasAlternatingBits2(n int) bool {
	for n != 0 {
		if (n & 1) == ((n >> 1) & 1) {
			return false // 相邻相等
		}
		n >>= 1
	}
	return true
}