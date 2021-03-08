package middle

import "math"

func divide(dividend int, divisor int) int {
	a := float64(dividend/divisor)
	if a < math.MinInt32 || a > math.MaxInt32{
		return math.MaxInt32
	}
	if a > 0{
		return int(math.Floor(a))
	}
	return int(math.Ceil(a))
}
