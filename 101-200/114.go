package _01_200



func flatten(root *TreeNode)  {
	if root == nil {
		return
	}
	flatten(root.Left)
	if root.Left != nil {
		l := root.Left
		for l.Right != nil {
			l = l.Right
		}
		l.Right = root.Right
		root.Right = root.Left
		root.Left = nil
	}
	flatten(root.Right)
}