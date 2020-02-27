package stack

type LinkStack struct {
	top *node
	len int
}

type node struct {
	value interface{}
	next  *node
}

func NewLinkStack() *LinkStack {
	return &LinkStack{}
}

func (this *LinkStack) Len() int {
	return this.len
}

//查看栈顶元素
func (this *LinkStack) Peek() interface{} {
	if this.top == nil {
		return nil
	}
	return this.top.value
}

//入栈
func (this *LinkStack) Push(value interface{}) {
	n := &node{
		value: value,
		next:  this.top,
	}
	this.top = n
	this.len++
}

//出栈
func (this *LinkStack) Pop() interface{} {
	if this.len == 0 {
		return nil
	}
	n := this.top
	this.top = this.top.next
	this.len--
	return n.value
}
