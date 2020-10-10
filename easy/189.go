package easy

import "fmt"

func rotate(nums []int, k int) {
	reverse(nums)
	reverse(nums[:k%len(nums)])
	reverse(nums[k%len(nums):])
}

func reverse(arr []int) {
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-i-1] = arr[len(arr)-i-1], arr[i]
	}
}

func Rotate2(nums []int, k int) {
	l := k % len(nums)
	nums = append(nums[len(nums)-l:], nums[:len(nums)-l]...)
	fmt.Println(nums)
}