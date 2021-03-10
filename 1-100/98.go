package __100

import "math"

/**
没通过
 */
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var order func(*TreeNode)
	cur := math.MinInt64
	ans := true
	order = func(root *TreeNode) {
		if root == nil {
			return
		}
		order(root.Left)
		if cur > root.Val {
			ans = false
			return
		}else {
			cur = root.Val
		}
		order(root.Right)
	}
	return ans
}











