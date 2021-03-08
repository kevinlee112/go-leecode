package _01_300

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	l, r := 0, 0
	lNode, rNode := root, root
	for lNode != nil {
		 l++
		 lNode = lNode.Left
	}
	for rNode != nil {
		r++
		rNode = rNode.Right
	}
	if l == r {
		return 1<<l - 1
	}

	return 1 + countNodes(root.Right) + countNodes(root.Left)
}
