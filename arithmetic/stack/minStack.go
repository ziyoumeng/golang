package stack

type MinStack struct {
	arr []int
	min []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		arr: make([]int, 0),
		min: make([]int, 0),
	}
}

func (this *MinStack) Push(x int) {
	this.arr = append(this.arr, x)
	if len(this.min) == 0{
		this.min = append(this.min, x)
	}else{
		min := this.min[len(this.min)-1]
		if x <  min{
			min  = x
		}
		this.min = append(this.min, min)
	}
}

func (this *MinStack) Pop() {
	if len(this.arr) > 0 {
		this.arr = this.arr[:len(this.arr)-1]
		this.min = this.min[:len(this.min)-1]
	}
}

func (this *MinStack) Top() int {
	if len(this.arr) > 0 {
		return this.arr[len(this.arr)-1]
	}
	return 0
}

func (this *MinStack) GetMin() int {
	if len(this.min) >0 {
		return this.min[len(this.min)-1]
	}
	return 0
}
