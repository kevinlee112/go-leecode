package easy

type TreeNode1 struct {
	Val   int
	Left  *TreeNode1
	Right *TreeNode1
}

var ress bool
var cur int

func isUnivalTree(root *TreeNode1) bool {
	if root == nil {
		return true
	}
	ress = true
	cur = root.Val
	dfs3(root)
	return ress
}

func dfs3(root *TreeNode1) {
	if root != nil {
		if root.Val != cur {
			ress = false
			return
		}
		dfs3(root.Left)
		dfs3(root.Right)
	}
}