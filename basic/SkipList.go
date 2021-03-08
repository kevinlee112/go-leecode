package basic

import (
"math/rand"
"sync"
"time"
)

//跳跃表节点
type SkipListNode struct {
	key  int
	data interface{}
	next []*SkipListNode //指针数组
}

//跳跃表链表
type SkipList struct {
	head   *SkipListNode
	tail   *SkipListNode
	length int           //链表长度
	level  int           //层数
	mut    *sync.RWMutex //互斥锁 保证线程安全
	rand   *rand.Rand    //随机数 满足P=0.5几何分布
}

//构造跳跃表
func newSkipList(level int) *SkipList {
	list := &SkipList{}
	if level <= 0 {
		return list
	}
	list.level = level
	list.head = &SkipListNode{next: make([]*SkipListNode, level, level)}
	list.tail = &SkipListNode{}
	list.mut = &sync.RWMutex{}
	list.rand = rand.New(rand.NewSource(time.Now().UnixNano()))
	for index := range list.head.next {
		list.head.next[index] = list.tail
	}

	return list
}

//随机生成器函数 满足P=0.5的几何分布
func (list *SkipList) randomLevel() int {
	level := 1
	for ; level < list.level && list.rand.Uint32()&0x1 == 1; level++ {
	}
	return level
}

//插入节点
func (list *SkipList) Add(key int, data interface{}) {

	list.mut.Lock()
	defer list.mut.Unlock()
	//1.确定插入深度
	level := list.randomLevel()
	//2.查找插入部位
	update := make([]*SkipListNode, level, level)
	node := list.head
	for index := level - 1; index >= 0; index-- {
		for {
			node1 := node.next[index]
			if node1 == list.tail || node1.key > key {
				update[index] = node //找到一个插入部位
				break
			} else if node1.key == key {
				node1.data = data
				return
			} else {
				node = node1
			}
		}
	}
	//3.执行插入
	newNode := &SkipListNode{key, data, make([]*SkipListNode, level, level)}
	for index, node := range update {
		node.next[index], newNode.next[index] = newNode, node.next[index]
	}
	list.length++
}

//删除节点
func (list *SkipList) Remove(key int) bool {
	list.mut.Lock()
	defer list.mut.Unlock()
	//1.查找删除节点
	node := list.head
	remove := make([]*SkipListNode, list.level, list.level)
	var target *SkipListNode
	for index := len(node.next) - 1; index >= 0; index-- {
		for {
			node1 := node.next[index]
			if node1 == list.tail || node1.key > key {
				break
			} else if node1.key == key {
				remove[index] = node //找到啦
				target = node1
				break
			} else {
				node = node1
			}
		}
	}

	//2.执行删除
	if target != nil {
		for index, node1 := range remove {
			if node1 != nil {
				node1.next[index] = target.next[index]
			}
		}
		list.length--
		return true
	}
	return false
}
//查找节点
func (list *SkipList) Find(key int) interface{} {
	list.mut.Lock()
	defer list.mut.Unlock()
	//1.查找删除节点
	node := list.head
	for index := len(node.next) - 1; index >= 0; index-- {
		for {
			node1 := node.next[index]
			if node1 == list.tail || node1.key > key {
				break
			} else if node1.key == key {
				return node1.data
			} else {
				node = node1
			}
		}
	}
	return nil
}
//获取链表长度
func (list *SkipList) Length() int {
	list.mut.RLock()
	defer list.mut.RUnlock()
	return list.length
}

