package __100

/**
恢复二叉搜索树
 */
func recoverTree(root *TreeNode) {
	var list []*TreeNode
	var recoverTreeHelper func(*TreeNode)
	recoverTreeHelper = func(root *TreeNode) {
		if root == nil{
			return
		}
		recoverTreeHelper(root.Left)
		list = append(list, root)
		recoverTreeHelper(root.Right)
	}
	recoverTreeHelper(root)
	x, y := -1, -1
	for i:=0;i<len(list)-1;i++{
		if list[i].Val > list[i+1].Val{
			y = i+1
			if x == -1 {
				x = i
			}
		}
	}
	list[x].Val, list[y].Val = list[y].Val, list[x].Val
}
