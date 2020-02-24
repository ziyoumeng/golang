package singleLink

import "fmt"


/*
输入: 1->2->3->4->5->NULL
输出: 5->4->3->2->1->NULL
 */
func reverseLink(head *LinkNode)*LinkNode{
	var prev *LinkNode
	cur := head
	for cur !=nil{
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev
}
func reverseLinkRecursion(head *LinkNode)*LinkNode{
	return recursion(nil, head)
}

func recursion(prev, cur *LinkNode)*LinkNode{
	if cur == nil{
		return prev
	}
	tmpNext := cur.Next
	cur.Next = prev
	return recursion(cur,tmpNext)
}

func main() {
	data := []int{1,2,3,4,5}
	head := makeLink(data)
	printLink(head)
	head = reverseLink(head)
	printLink(head)
	head = reverseLinkRecursion(head)
	printLink(head)
}

func makeLink(data []int)*LinkNode{
	var head, prev *LinkNode
	for i, val := range data{
		cur := &LinkNode{Val: val}
		if i== 0{
			head = cur
		}else{
			prev.Next = cur
		}
		prev = cur
	}
	return  head
}

func printLink(head *LinkNode){
	cur := head
	for cur!=nil{
		fmt.Printf("%d->",cur.Val)
		cur = cur.Next
	}
	fmt.Println("")
}
