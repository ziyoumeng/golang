package linked_list

import "fmt"

type Node struct {
	value, key interface{}
	next       *Node
	prev       *Node
}

func NewNode(key, value interface{}) *Node {
	return &Node{
		value: value,
		key:   key,
	}
}

func (n Node) GetValue() interface{} {
	return n.value
}

func (n Node) GetKey() interface{} {
	return n.key
}

//双向循环链表
type DoublyLinkedList struct {
	sentinel *Node //哨兵
}

func NewDoublyLinkedList() *DoublyLinkedList {
	sentinel := &Node{}
	sentinel.next = sentinel
	sentinel.prev = sentinel
	return &DoublyLinkedList{sentinel: sentinel}
}

func (s *DoublyLinkedList) Search(key interface{}) *Node {
	if key == nil {
		return nil
	}
	for cur := s.sentinel.next; cur != s.sentinel; cur = cur.next {
		if cur.key == key {
			return cur
		}
	}
	return nil
}

//插入头部
func (s *DoublyLinkedList) Prepend(x *Node) {
	if x == nil || x.key == nil {
		return
	}
	head := s.sentinel.next
	x.next = head
	head.prev = x

	s.sentinel.next = x
	x.prev = s.sentinel
}

//注意：x必须是链表中的结点
func (s *DoublyLinkedList) Delete(x *Node) {
	if x == nil {
		return
	}
	x.prev.next = x.next
	x.next.prev = x.prev
}

func (s *DoublyLinkedList) Print() {
	for cur := s.sentinel.next; cur != s.sentinel; cur = cur.next {
		fmt.Printf("%+v-->", cur.key)
	}
	fmt.Println()
}

func (s *DoublyLinkedList) GetTail() *Node {
	if s.sentinel.prev == s.sentinel {
		return nil
	}
	return s.sentinel.prev
}

func (s *DoublyLinkedList) IsHead(x *Node) bool {
	return s.sentinel.next == x
}
