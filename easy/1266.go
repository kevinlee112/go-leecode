package easy

import "math"

func minTimeToVisitAllPoints(points [][]int) int {
	res := 0
	l := len(points)
	for i:=1;i<l;i++ {
		res += dsBt(points[i-1], points[i])
	}
	return res
}

func dsBt(n, m []int) int {
	res := math.Abs(float64(m[0] - n[0]))
	if math.Abs(float64(m[1] - n[1])) > res {
		res = math.Abs(float64(m[1] - n[1]))
	}
	return int(res)
}