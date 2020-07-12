package easy

func DistributeCandies(candies int, num_people int) []int {
	res := make([]int, num_people)
	n := 0
	for candies >0 {
		n++
		if n <= candies {
			res[(n-1)%num_people] += n
			candies -= n
		}else {
			res[(n-1)%num_people] += candies
			candies = 0
		}
	}
	return res
}