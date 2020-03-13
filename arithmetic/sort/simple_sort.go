package sort

import "fmt"

type sort interface {
	Len() int
}

//冒泡排序
//冒泡和选择排序相同点：保持前k大元素已排序
func BubbleSort(arr []int) {
	var hasSwap bool
	for i := 1; i < len(arr); i++ { //第几轮
		for j := 0; j < len(arr)-i; j++ { //每轮冒泡一个元素
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				hasSwap = true
			}
		}
		if !hasSwap {
			break
		}
	}
}

//插入排序
//保持局部有序然后扩大到整体
func InsertionSort(arr []int) {
	for i:=1; i < len(arr); i++ {
		val:= arr[i]
		j:=i-1
		for ; j >= 0;j-- {
			if val<arr[j] {
				arr[j+1]=arr[j] //比冒泡排序少一次赋值
			}else{
				break
			}
		}
		arr[j+1] = val
	}
}


//希尔排序
func ShellSort(arr []int) {
	for step := len(arr) / 2; step > 0; step /= 2 {
		for i:=step;i<len(arr); i++ {
			j:= i
			for j-step >=0 && arr[j]<arr[j-step]{
				arr[j],arr[j-step] = arr[j-step],arr[j]
				j -= step
			}
		}
		//printArray(arr,step)
	}
}

//选择排序
func SelectionSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		minIndex := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		if minIndex != i {
			arr[i], arr[minIndex] = arr[minIndex], arr[i]
		}
	}
}

func printArray(arr []int,step int){
	for i := 0; i < len(arr); i++{
		if i!=0 && i%step==0{
			fmt.Println()
		}
		fmt.Printf("%d ",arr[i])
	}
	fmt.Println("\n----")
}