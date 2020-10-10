package middle

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	res := append(preorderTraversal(root.Left), preorderTraversal(root.Right)...)
	res = append([]int{root.Val}, res...)
	return res
}