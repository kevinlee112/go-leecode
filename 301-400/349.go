package _01_400

func intersection(nums1 []int, nums2 []int) []int {
	res := make([]int, 0)
	if len(nums1) == 0 || len(nums2) == 0 {
		return res
	}
	m := make(map[int]bool)
	for _, v := range nums1{
		m[v] = true
	}

	for _, v := range nums2{
		if m[v] == true {
			res = append(res, v)
			m[v] = false
		}
	}
	return  res
}
