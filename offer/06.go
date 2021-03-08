package offer

type ListNode struct {
	Val  int
	Next *ListNode
}


func reversePrint(head *ListNode) (res []int) {
	for head != nil {
		temp := head
		defer func() {
			res = append(res, temp.Val)
		}()
		head = head.Next
	}
	return
}

