package _01_400

import "container/heap"

/*
给定一个非空的整数数组，返回其中出现频率前 k 高的元素。

 

示例 1:

输入: nums = [1,1,1,2,2,3], k = 2
输出: [1,2]
示例 2:

输入: nums = [1], k = 1
输出: [1]
 */

func topKFrequent(nums []int, k int) []int {
	if k == 0 || len(nums) == 0 {
		return make([]int, 0)
	}
	// 1. 初始化map
	m := make(map[int]int)
	for _, v := range nums {
		m[v] = m[v] + 1
	}

	// 2. 放到小根堆里面
	h := &NodeHeap{}
	topK := min(k, len(m))
	size := 0
	for k, v := range m {
		if size < topK {
			heap.Push(h, &Node{
				val:   k,
				times: v,
			})
			size++
		} else {
			if v > (*h)[0].times {
				heap.Pop(h)
				heap.Push(h, &Node{
					val:   k,
					times: v,
				})
			}
		}

	}

	// 3.收集答案
	res := make([]int, 0, topK)
	for i := 0; i < topK; i++ {
		res = append(res, heap.Pop(h).(*Node).val)
	}
	return res
}

type Node struct {
	val   int
	times int
}

type NodeHeap []*Node

func (h NodeHeap) Len() int { return len(h) }

// 小根堆
func (h NodeHeap) Less(i, j int) bool { return h[i].times < h[j].times }

func (h NodeHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *NodeHeap) Push(x interface{}) {
	*h = append(*h, x.(*Node))
}

func (h *NodeHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}