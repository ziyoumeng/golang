package stack

type SequentialStack struct {
	arr []interface{}
}

func NewSequentialStack() *SequentialStack {
	return &SequentialStack{
		arr: make([]interface{}, 0),
	}
}

func (this *SequentialStack) Push(x interface{}) {
	this.arr = append(this.arr, x)
}

func (this *SequentialStack) Pop() interface{} {
	if len(this.arr) == 0 {
		return nil
	}
	top := this.arr[len(this.arr)-1]
	this.arr = this.arr[:len(this.arr)-1]
	return top
}

func (this *SequentialStack) Peek() interface{} {
	if len(this.arr) == 0 {
		return nil
	}
	return this.arr[len(this.arr)-1]
}

func (this *SequentialStack) Len() int {
	return len(this.arr)
}
