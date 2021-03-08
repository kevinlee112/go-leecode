package _01_200

/**
给定一个二叉树，返回其节点值的锯齿形层序遍历。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。

例如：
给定二叉树 [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回锯齿形层序遍历如下：

[
  [3],
  [20,9],
  [15,7]
]
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	result := make([][]int, 0)
	queue := []*TreeNode{root}
	queue = append(queue, root)
	//反转标志位
	flag := false
	for len(queue) > 0 {
		list := make([]int, 0)
		//当前层的个数
		l := len(queue)
		for i:=-0;i<l;i++ {
			level := queue[0]
			queue = queue[1:]
			list = append(list, level.Val)
			if level.Left != nil {
				queue = append(queue, level.Left)
			}
			if level.Right != nil {
				queue = append(queue, level.Right)
			}
		}
		if(flag){
			reverse(list)
		}
		result = append(result, list)
		// 将标志位取反
		flag = !flag
	}
	return result
}
func reverse(nums []int) {
	for i := 0; i < len(nums) / 2; i++ {
		nums[i], nums[len(nums)-1-i] = nums[len(nums)-1-i], nums[i]
	}
}
