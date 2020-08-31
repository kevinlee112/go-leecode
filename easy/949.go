package easy

import (
	"fmt"
	"math"
)

var result string // 最终结果
var maxTime int // 最大时间
func largestTimeFromDigits(A []int) string {
	seen, selected := make(map[int]bool), []int{}
	result, maxTime = "", 0
	backTrack(seen, selected, A)
	return result
}

func backTrack(seen map[int]bool, selected []int, A []int)  {
	if len(selected) == 4 {
		t := makeTimeValue(selected)
		if t >= maxTime {
			maxTime = t
			result = makeTime(selected)
		}
	}
	for k, v := range A {
		if seen[k] {
			continue
		}
		selected = append(selected, v)
		// 往后组成的时间肯定格式不对直接剪枝
		if !isValidTime(selected) {
			selected = selected[:len(selected)-1]
			continue
		}
		seen[k] = true
		backTrack(seen, selected, A)
		seen[k] = false
		selected = selected[:len(selected)-1]
	}
}

// 是否是合法的时间
func isValidTime(t []int) bool {
	if makeTimeValue(t) > 2359 {
		return false
	}
	if len(t) >= 3 && t[2] >= 6 {
		return false
	}
	return true
}

// 时间转为数值
func makeTimeValue(t []int) (result int) {
	for _, v := range t {
		result = result * 10 +v
	}
	result *= int(math.Pow(10, float64(4-len(t))))
	return
}

// 时间转为字符串格式
func makeTime(t []int) string {
	return fmt.Sprintf("%d%d:%d%d", t[0], t[1], t[2], t[3])
}
