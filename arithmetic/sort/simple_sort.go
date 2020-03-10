package sort

type sort interface{
	Len()int
}

//冒泡排序
func BubbleSort(arr []int){
	var hasSwap bool
	for i:=1; i < len(arr); i++{ //第几轮
		for j:=0;j<len(arr)-i;j++{ //每轮冒泡一个元素
			if arr[j] >arr[j+1]{
				arr[j],arr[j+1] = arr[j+1],arr[j]
				hasSwap = true
			}
		}
		if !hasSwap{
			break
		}
	}
}

func InsertionSort(arr []int) {
	for i:=1; i < len(arr); i++ {
		val:= arr[i]
		j:=i-1
		for ; j >= 0;j-- {
			if val<arr[j] {
				arr[j+1]=arr[j]
			}else{
				break
			}
		}
		arr[j+1] = val
	}
}
