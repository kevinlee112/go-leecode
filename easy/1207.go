package easy


func uniqueOccurrences(arr []int) bool {
	dict := make(map[int]int)
	for _, v := range arr{
		dict[v]++
	}
	dictCount := make(map[int]bool)
	for _,v := range dict{
		if dictCount[v] {
			return false
		}
		dictCount[v] = true
	}
	return true
}
