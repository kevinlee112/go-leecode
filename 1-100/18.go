package __100

import (
	"sort"
)

func FourSum(nums []int, target int) [][]int {
	var ans [][]int
	sort.Ints(nums)
	for a:=0;a<len(nums)-3;a++ {
		if a > 0 && nums[a] == nums[a-1] {
			continue
		}
		for b:=a+1;b<len(nums)-2;b++ {
			if b>a+1&&nums[b]==nums[b-1] {
				continue
			}
			p1, p2 := b+1, len(nums)-1
			for p1 < p2 {
				sum := nums[a] + nums[b] + nums[p1] + nums[p2]
				if sum == target {
					ans = append(ans, []int{nums[a], nums[b], nums[p1], nums[p2]})
					for p1 < p2 && nums[p1] == nums[p1+1]{
						p1++
					}
					for p1 < p2 && nums[p2] == nums[p2-1]{
						p2--
					}
					p1++
					p2--
				}else if sum < target {
					p1++
				}else {
					p2--
				}
			}
		}
	}
	return ans
}
