package sort

import (
	"fmt"
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T){
	cases := []struct {
		arr []int
		want []int
	}{
		{nil,nil},
		{[]int{1},[]int{1}},
		{[]int{2,1},[]int{1,2}},
		{[]int{4,3,2,1},[]int{1,2,3,4}},
	}
	for _,c :=range cases{
		t.Run(fmt.Sprintf("length is %d",len(c.arr)),func(t *testing.T) {
			BubbleSort(c.arr)
			if !reflect.DeepEqual(c.arr,c.want){
				t.Errorf("output:%+v, want:%+v",c.arr,c.want)
			}
		})
	}
}

func TestInsertionSort(t *testing.T) {
	cases := []struct {
		arr []int
		want []int
	}{
		{nil,nil},
		{[]int{1},[]int{1}},
		{[]int{2,1},[]int{1,2}},
		{[]int{4,3,2,1},[]int{1,2,3,4}},
	}
	for _,c :=range cases{
		t.Run(fmt.Sprintf("length is %d",len(c.arr)),func(t *testing.T) {
			InsertionSort(c.arr)
			if !reflect.DeepEqual(c.arr,c.want){
				t.Errorf("output:%+v, want:%+v",c.arr,c.want)
			}
		})
	}
}

func TestSelectionSort(t *testing.T) {
	cases := []struct {
		arr []int
		want []int
	}{
		{nil,nil},
		{[]int{1},[]int{1}},
		{[]int{2,1},[]int{1,2}},
		{[]int{4,3,2,1},[]int{1,2,3,4}},
	}
	for _,c :=range cases{
		t.Run(fmt.Sprintf("length is %d",len(c.arr)),func(t *testing.T) {
			SelectionSort(c.arr)
			if !reflect.DeepEqual(c.arr,c.want){
				t.Errorf("output:%+v, want:%+v",c.arr,c.want)
			}
		})
	}
}


func TestShellSort(t *testing.T) {
	cases := []struct {
		arr []int
		want []int
	}{
		{nil,nil},
		{[]int{1},[]int{1}},
		{[]int{2,1},[]int{1,2}},
		{[]int{4,3,2,1},[]int{1,2,3,4}},
		{[]int{9,8,7,6,5,4,3,2,1,0},[]int{0,1,2,3,4,5,6,7,8,9}},
	}
	for _,c :=range cases{
		t.Run(fmt.Sprintf("length is %d",len(c.arr)),func(t *testing.T) {
			ShellSort(c.arr)
			if !reflect.DeepEqual(c.arr,c.want){
				t.Errorf("output:%+v, want:%+v",c.arr,c.want)
			}
		})
	}
}

func TestMergeSort(t *testing.T) {
	cases := []struct {
		arr []int
		want []int
	}{
		{nil,nil},
		{[]int{1},[]int{1}},
		{[]int{2,1},[]int{1,2}},
		{[]int{4,3,2,1},[]int{1,2,3,4}},
		{[]int{9,8,7,6,5,4,3,2,1,0},[]int{0,1,2,3,4,5,6,7,8,9}},
	}
	for _,c :=range cases{
		t.Run(fmt.Sprintf("length is %d",len(c.arr)),func(t *testing.T) {
			MergeSort(c.arr)
			if !reflect.DeepEqual(c.arr,c.want){
				t.Errorf("output:%+v, want:%+v",c.arr,c.want)
			}
		})
	}
}

func TestQuickSort(t *testing.T) {
	cases := []struct {
		arr []int
		want []int
	}{
		{nil,nil},
		{[]int{1},[]int{1}},
		{[]int{2,1},[]int{1,2}},
		{[]int{4,3,2,1},[]int{1,2,3,4}},
		{[]int{9,8,7,6,5,4,3,2,1,0},[]int{0,1,2,3,4,5,6,7,8,9}},
	}
	for _,c :=range cases{
		t.Run(fmt.Sprintf("length is %d",len(c.arr)),func(t *testing.T) {
			QuickSort(c.arr)
			if !reflect.DeepEqual(c.arr,c.want){
				t.Errorf("output:%+v, want:%+v",c.arr,c.want)
			}
		})
	}
}

func TestFindNthElement(t *testing.T) {
	cases := []struct {
		arr []int
		nth  int
		want int
	}{
		{nil,0,-1},
		{[]int{1},1,1},
		{[]int{2,1},2,1},
		{[]int{4,3,2,1},3,2},
		{[]int{9,5,4,3,8,7,6,2,1,0,10},5,6},
	}
	for _,c :=range cases{
		t.Run(fmt.Sprintf("length is %d",len(c.arr)),func(t *testing.T) {
			result := FindNthElement(c.arr,c.nth)
			if !reflect.DeepEqual(result,c.want){
				t.Errorf("output:%d, want:%+v",result,c.want)
			}
		})
	}
}
