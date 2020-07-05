package easy

import "sort"

func SortByBits(arr []int) []int {
	sort.Slice(arr, func(i, j int) bool {
		if OneNum(arr[i]) < OneNum(arr[j]) {
			return true
		} else if OneNum(arr[i]) == OneNum(arr[j]) && arr[i] <= arr[j] {
			return true
		}
		return false
	})
	return arr
}

func OneNum(num int) int {
	count := 0
	for num != 0 {
		if num&1 == 1 {
			count++
		}
		num = num >> 1
	}
	return count
}