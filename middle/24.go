package middle

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	curr, pre := head, head.Next
	pre.Next = curr
	curr.Next = swapPairs(pre.Next)
	return pre
}
