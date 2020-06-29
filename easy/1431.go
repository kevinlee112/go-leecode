package easy


func KidsWithCandies(candies []int, extraCandies int) []bool {

	n := len(candies)
	maxCandies := 0
	for _, v := range candies {
		maxCandies = max(v, maxCandies)
	}
	ret := make([]bool, n)
	for i:=0;i<n;i++ {
		ret[i] = candies[i] + extraCandies > maxCandies
	}

	return ret
}


func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
