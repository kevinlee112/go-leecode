package easy

func FindLucky(arr []int) int {
	res := -1
	dict := make(map[int]int)
	for _,v := range arr{
		dict[v]++
	}
	for k,v := range dict{
		if k == v && res < v {
			res = v
		}
	}

	return res
}