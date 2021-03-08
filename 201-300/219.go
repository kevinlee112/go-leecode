package _01_300







func containsNearbyDuplicate(nums []int, k int) bool {
	h := make(map[int]int)
	var indexArr []int
	for index, v := range nums {
		_, ok := h[v]
		if ok {
			return true
		}
		h[v] = index
		indexArr = append(indexArr, index)
		if len(indexArr) > k {
			delete(h, nums[indexArr[0]])
			indexArr = indexArr[1:]
		}
	}
	return false
}