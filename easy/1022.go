package easy

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var total int
func sumRootToLeaf(root *TreeNode) int {
	total = 0
	dfs2(root, 0)
	return total
}

func dfs2(root *TreeNode, num int) {
	if root == nil {
		return
	}
	num = num << 1 + root.Val
	if root.Left == nil && root.Right == nil {
		total += num
	}
	dfs2(root.Left, num)
	dfs2(root.Right, num)
}
