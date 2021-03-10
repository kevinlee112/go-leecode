package __100
/*
给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。

你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。



示例 1：
输入：head = [1,2,3,4]
输出：[2,1,4,3]
示例 2：

输入：head = []
输出：[]
示例 3：

输入：head = [1]
输出：[1]

*/

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next==nil {
		return head
	}
	curr, pre := head, head.Next
	curr.Next = swapPairs(pre.Next)
	pre.Next = curr
	return pre
}