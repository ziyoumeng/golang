package sort

import (
	"math/rand"
)

func QuickSort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	p1, p2 := partion(arr)
	QuickSort(p1)
	QuickSort(p2)
}
func getPivotIndex(length int) int {
	return rand.Intn(length)
	//return length - 1
}

func partion(arr []int) (p1, p2 []int) {
	pIndex := getPivotIndex(len(arr))

	//将pivot元素交换到数组尾部作为哨兵
	lastIndex := len(arr) - 1
	if pIndex != lastIndex {
		arr[lastIndex], arr[pIndex] = arr[pIndex], arr[lastIndex]
	}

	pivot := arr[lastIndex]
	i, j := 0, 0
	for ; j < len(arr); j++ { //[0,i)比pivot小，[i,j)比pivot大,[j,lastIndex)待划分
		if arr[j] < pivot{
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[lastIndex] = arr[lastIndex], arr[i]
	return arr[:i],arr[i+1:]
}
//找第k大元素 O(n)
func FindNthElement(arr []int,k int)int{
	if len(arr) == 0{
		return -1
	}
	pIndex := getPivotIndex(len(arr))

	//将pivot元素交换到数组尾部作为哨兵
	lastIndex := len(arr) - 1
	if pIndex != lastIndex {
		arr[lastIndex], arr[pIndex] = arr[pIndex], arr[lastIndex]
	}

	pivot := arr[lastIndex]
	i, j := 0, 0
	for ; j < len(arr); j++ {
		if arr[j] > pivot{
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[lastIndex] = arr[lastIndex], arr[i]
	//fmt.Println(arr,i,k)
	if i > k-1 {
		return FindNthElement(arr[:i],k)
	}else if i<k-1{
		return FindNthElement(arr[i+1:],k-i-1)
	}else {
		return arr[i]
	}
	return -1
}
