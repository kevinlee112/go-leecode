package basic

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//前序遍历
func PreOrderTraversal(tree *TreeNode)  {
	if tree == nil {
		return
	}

	fmt.Println("%d ->", tree.Val)
	PreOrderTraversal(tree.Left)
	PreOrderTraversal(tree.Right)
}
// 中序遍历
func MidOrderTraversal(tree *TreeNode) {
	if tree == nil {
		return
	}

	MidOrderTraversal(tree.Left)
	fmt.Printf(" %d -> ", tree.Val)
	MidOrderTraversal(tree.Right)
}
// 后序遍历
func PostOrderTraversal(tree *TreeNode) {
	if tree == nil {
		return
	}

	PostOrderTraversal(tree.Left)
	PostOrderTraversal(tree.Right)
	fmt.Printf(" %d -> ", tree.Val)
}

// 按层遍历

func LevelOrderTraversal(tree *TreeNode) {
	if tree == nil {
		return
	}

	// 采用队列实现
	queue := make([]*TreeNode, 0)
	queue = append(queue, tree) // queue push
	for len(queue) > 0 {
		tree = queue[0]
		fmt.Printf(" %d -> ", tree.Val)

		queue = queue[1:] // queue pop

		if tree.Left != nil {
			queue = append(queue, tree.Left)
		}
		if tree.Right != nil {
			queue = append(queue, tree.Right)
		}
	}
}