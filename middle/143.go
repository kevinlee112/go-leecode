package middle

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil{
		return
	}
	//查找last
	last := head
	for last.Next.Next != nil {
		last = last.Next
	}

	tmp := last.Next
	last.Next = nil
	pre := head.Next
	head.Next = tmp
	head.Next.Next = pre
	reorderList(head.Next.Next)
}


