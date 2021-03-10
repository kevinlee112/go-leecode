package __100

/**
反转从位置 m 到 n 的链表。请使用一趟扫描完成反转。
说明:
1 ≤ m ≤ n ≤ 链表长度。

示例:
输入: 1->2->3->4->5->NULL, m = 2, n = 4
输出: 1->4->3->2->5->NULL
 */
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	if m == 1 {
		return reverseN(head, n)
	}
	head.Next = reverseBetween(head.Next, m-1, n-1)
	return head
}

var succ *ListNode


func reverseN(head *ListNode, n int) *ListNode {
	if n == 1 {
		succ = head.Next
		return head
	}
	rev := reverseN(head.Next, n-1)
	head.Next.Next = head
	head.Next = succ
	return rev
}







