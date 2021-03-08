package _01_600

/**

 */

func NextGreaterElements(nums []int) []int {
	length := len(nums)
	if length <= 0 {
		return []int{}
	}
	res:= make([]int, length)
	stack := make([]int,0)
	for i:=0; i< length; i++{
		res[i] = -1
		for len(stack) >0 && nums[stack[len(stack)-1]] < nums[i] {
			res[stack[len(stack)-1]] = nums[i]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	for len(stack)>1 {
		if nums[stack[len(stack)-1]] < nums[stack[0]] {
			res[stack[len(stack)-1]] = nums[stack[0]]
			stack = stack[:len(stack)-1]
		}else {
			break
		}
	}

	return res
}