package hard

func reverseKGroup(head *ListNode, k int) *ListNode {
	cur := head
	for i := 1; i < k && cur != nil; i++ {
		cur = cur.Next
	}
	if cur == nil {
		return head
	}

	next := cur.Next
	cur.Next = nil
	cur = reverseList(head)
	head.Next = reverseKGroup(next, k)

	return cur
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	left, right := head, head.Next
	left.Next = nil
	for right != nil {
		next := right.Next
		right.Next = left
		left, right = right, next
	}
	return head
}