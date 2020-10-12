package middle

import "sort"

func fourSum(nums []int, target int) [][]int {
	ans := make([][]int, 0)
	n := len(nums)
	if n < 4 {
		return ans
	}

	sort.Ints(nums)

	for a := 0; a <= n-4; a++ {
		if a>0 && nums[a] == nums[a-1] { // 剪枝
			continue
		}
		for b := a+1; b <= n-3; b++ {
			if b > a+1 && nums[b] == nums[b-1] { // 剪枝
				continue
			}
			c, d := b+1, n-1
			for c<d {
				sum := nums[a]+nums[b]+nums[c]+nums[d]
				if sum == target{
					ans = append(ans, []int{nums[a], nums[b], nums[c], nums[d]})
					for c<d && nums[c] == nums[c+1] {
						c++
					}
					for c<d && nums[d] == nums[d-1] {
						d--
					}
					c++
					d--
				} else if sum < target {
					c++
				} else {
					d--
				}
			}
		}
	}
	return ans
}
















