package middle

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	res := append(postorderTraversal(root.Left), postorderTraversal(root.Right)...)
	res = append(res, root.Val)
	return res
}