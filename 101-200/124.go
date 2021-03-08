package _01_200

/**
124. 二叉树中的最大路径和
给定一个非空二叉树，返回其最大路径和。

本题中，路径被定义为一条从树中任意节点出发，沿父节点-子节点连接，达到任意节点的序列。该路径至少包含一个节点，且不一定经过根节点。



示例 1：

输入：[1,2,3]

       1
      / \
     2   3

输出：6
示例 2：

输入：[-10,9,20,null,null,15,7]

   -10
   / \
  9  20
    /  \
   15   7

输出：42
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	maxSum := -1 << 31
	maxGain(root, &maxSum)
	return maxSum
}

func maxGain(tn *TreeNode, maxSum *int) int {
	if tn == nil {  //叶子节点
		return 0
	}
	left := maxGain(tn.Left, maxSum)
	right := maxGain(tn.Right, maxSum)
	if left < 0 {
		left = 0    //节点值小于0，则不参与路径
	}
	if right < 0 {
		right = 0
	}
	if tn.Val + left + right > *maxSum {
		*maxSum = tn.Val + left + right
	}

	//当前节点的路径最大值
	if left > right {
		return tn.Val + left
	}

	return tn.Val + right
}












