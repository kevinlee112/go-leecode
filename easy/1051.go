package easy

func heightChecker(heights []int) int {
	l := len(heights)
	dict := make(map[int]bool)
	min := heights[0]
	index := 0
	for i:=0;i<l;i++{
		min = heights[i]
		index = i
		for j:=l-1;j>i;j-- {
			if heights[j] < min {
				min = heights[j]
				index = j

			}
		}
		if index != i {
			heights[i], heights[index] = heights[index], heights[i]
			dict[i] = true
			dict[index] = true
		}
	}
	return len(dict)
}