package __100

/*
将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。 

示例：

输入：1->2->4, 1->3->4
输出：1->1->2->3->4->4
 */

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := new(ListNode)
	prev := head
	for l1!=nil && l2!=nil {
		if l1.Val<l2.Val {
			head.Next = l1
			l1 = l1.Next
		}else {
			head.Next = l2
			l2 = l2.Next
		}
		head = head.Next
	}
	if l1 !=nil {
		head.Next = l1
	}else {
		head.Next=l2
	}

	return prev.Next
}
