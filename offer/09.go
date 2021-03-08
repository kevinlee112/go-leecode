package offer

type CQueue struct {
	// 栈 A，用于添加元素
	stackA []int
	// 栈 B，用于取出元素
	stackB []int
}

func Constructor() CQueue {
	return CQueue{}
}

func (this *CQueue) AppendTail(value int)  {
	this.stackA = append(this.stackA, value)
}

func (this *CQueue) DeleteHead() int {
	if len(this.stackB) == 0 {
		if len(this.stackA) == 0 {
			return -1
		}
		for len(this.stackA) > 0 {
			index := len(this.stackA) - 1
			value := this.stackA[index]
			this.stackB = append(this.stackB, value)
			this.stackA = this.stackA[:index]
		}
	}
	index := len(this.stackB) - 1
	value := this.stackB[index]
	this.stackB = this.stackB[:index]

	return value
}
