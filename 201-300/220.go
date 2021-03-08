package _01_300

import (
	"math"
)

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	for i := range nums{
		for j:=i+1;j<=i+k && j <len(nums);j++{
			if math.Abs(float64(nums[i]-nums[j])) <= float64(t){
				return true
			}
		}
	}
	return false
}









