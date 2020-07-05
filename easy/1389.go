package easy

import "fmt"

func CreateTargetArray(nums []int, index []int) []int {
	res := make([]int, 0)
	for i, v := range index{
		rear := append([]int{}, res[v:]...)
		res = append(res[0:v], nums[i])
		res = append(res, rear...)
		fmt.Println(res)
	}
	return res
}
