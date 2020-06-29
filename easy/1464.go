package easy

func MaxProduct(nums []int) int {
	top1, top2 := 0, 0
	for _, v := range nums {
		if v >= top1 {
			top2 = top1
			top1 = v
		} else if v > top2 {
			top2 = v
		}
	}

	return (top1 - 1) * (top2 - 1)
}
