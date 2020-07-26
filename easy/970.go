package easy

func powerfulIntegers(x int, y int, bound int) []int {
	res := map[int]int{}

	helper(0, 0, 1, 1, x, y, bound, res)

	ans := []int{}
	for i,_ := range res {
		ans = append(ans, i)
	}
	return ans
}


func helper(i, j, numx, numy int, x, y int, bound int, res map[int]int) {
	num := numx + numy
	if num > bound {
		return
	}
	res[num] = 1

	if y != 1 {
		helper(i, j+1, numx, numy*y, x, y, bound, res)
	}

	if x != 1 {
		helper(i+1, j, numx*x, numy, x, y, bound, res)
	}
}