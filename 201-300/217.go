package _01_300


func containsDuplicate(nums []int) bool {
	m := make(map[int]bool)
	for _, v := range nums {
		if _, ok := m[v]; ok {
			return true
		} else {
			m[v] = true
		}
	}
	return false
}