package easy

import (
	"fmt"
	"math"
	"sort"
)

func MinimumAbsDifference(arr []int) [][]int {
	sort.Ints(arr)
	res := make([][]int, 0)
	min := arr[1] - arr[0]
	for i:=1;i<len(arr);i++ {
		fmt.Println(min)
		if (arr[i] - arr[i-1]) == min {
			res = append(res, arr[i-1:i+1])
		}else if math.Abs(float64(arr[i] - arr[i-1])) < float64(min) {
			min = arr[i] - arr[i-1]
			res = make([][]int, 0)
			res = append(res, arr[i-1:i+1])
		}
	}
	return res
}