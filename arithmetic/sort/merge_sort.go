package sort

func MergeSort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	q := len(arr) / 2
	MergeSort(arr[:q])
	MergeSort(arr[q:])
	merge_sort(arr,arr[:q],arr[q:])
}

func merge_sort(arr, part1, part2 []int){
	tmp := make([]int, 0, len(part1)+len(part2))
	i, j := 0, 0
	for i < len(part1) || j < len(part2) {
		if i < len(part1) && j < len(part2) {
			if part1[i] <= part2[j] {
				tmp = append(tmp, part1[i])
				i++
			} else {
				tmp = append(tmp, part2[j])
				j++
			}
			continue
		} else if i < len(part1) {
			tmp = append(tmp, part1[i:]...)
		} else if j < len(part2) {
			tmp = append(tmp, part2[j:]...)
		}
		break
	}

	for i := range tmp {
		arr[i] = tmp[i]
	}
}
