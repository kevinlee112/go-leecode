package middle

func maxArea(height []int) int {
	res, s := 0, 0
	for i, j := 0, len(height)-1;i != j; {
		if height[i] > height[j] {
			s = (j - i) * height[j]
			j--
		} else {
			s = (j - i) * height[i]
			i++
		}
		if s > res {
			res = s
		}
	}
	return res
}
