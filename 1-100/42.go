package __100

func trap(height []int)int{
	l := len(height)
	res := 0
	for i:=1;i<l-1;i++{
		leftMax := 0
		for j:=0;j<=i;j++ {
			leftMax = max(leftMax, height[j])
		}
		rightMax := 0
		for k:=i;k<l;k++ {
			rightMax = max(rightMax, height[k])
		}
		minVal := min(leftMax,rightMax)
		water := minVal-height[i]
		res =res+water
	}
	return res
}

func max(i, j int) int {
	if i < j {
		return j
	}
	return i
}
