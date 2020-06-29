package easy

func BuildArray(target []int, n int) []string {

	res := make([]string, 0)
	num := 0
	for i := 1; i <= n; i++ {
		if i > target[len(target)-1] {
			return res
		} else if i == target[num]  {
			res = append(res, "Push")
			num++
		} else {
			res = append(res, "Push")
			res = append(res, "Pop")
		}
	}
	return res
}
