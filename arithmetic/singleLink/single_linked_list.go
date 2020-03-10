package singleLink

type ListNode = LinkNode
//单链表删除倒数第n个元素
//1、题目已经保证n是合法的
//2、涉及删除 添加哨兵简化逻辑
//3、双指针实现一次遍历
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	guard := &ListNode{
		Next:head,
	}
	fast,slow := guard,guard
	for n >0 {
		fast = fast.Next
		n--
	}
	for fast.Next != nil{
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next  = slow.Next.Next
	return guard.Next
}
//876 找中间结点
//给定一个带有头结点 head 的非空单链表，返回链表的中间结点。
//如果有两个中间结点，则返回第二个中间结点。
func middleNode(head *ListNode) *ListNode {
	slow,fast := head,head
	for fast.Next!=nil &&fast.Next.Next !=nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	if fast.Next ==nil {
		return slow
	}else{
		return slow.Next
	}
}
