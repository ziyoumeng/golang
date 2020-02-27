package LRU

import (
	"mytest/linked_list"
	"sync"
)

type Node = linked_list.Node

type LRU struct {
	count      int
	maxNum     int
	m          map[interface{}]*Node
	doublyList *linked_list.DoublyLinkedList

	sync.RWMutex
}

func NewLRU(maxNum int) *LRU {
	return &LRU{
		maxNum:     maxNum,
		m:          make(map[interface{}]*Node),
		doublyList: linked_list.NewDoublyLinkedList(),
	}
}

func (l *LRU) Put(key, value interface{}) {
	l.Lock()
	defer l.Unlock()

	if _, exist := l.m[key]; exist {
		return
	}

	if l.count == l.maxNum {
		//小的优化点 重复利用删除的结点
		tail := l.doublyList.GetTail()
		l.doublyList.Delete(tail)
		delete(l.m, tail.GetKey())
	} else {
		l.count++
	}

	node := linked_list.NewNode(key, value)
	l.doublyList.Prepend(node)
	l.m[key] = node
}

func (l *LRU) Get(key interface{}) (interface{}, bool) {
	l.Lock()
	defer l.Unlock()

	if node, exist := l.m[key]; exist {
		l.moveToHead(node)
		return node.GetValue(), exist
	} else {
		return nil, exist
	}
}

func (l *LRU) Print() {
	l.doublyList.Print()
}

func (l *LRU) moveToHead(node *Node) {
	if l.doublyList.IsHead(node) {
		return
	}
	l.doublyList.Delete(node)
	l.doublyList.Prepend(node)
}
