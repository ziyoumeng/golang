package sort

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T){
	arrs :=[]int{5,4,3,2,1}
	BubbleSort(arrs)
	expected:=[]int{1,2,3,4,5}
	if !reflect.DeepEqual(arrs,expected){
		t.Errorf("not equal")
	}
}

func TestInsertionSort(t *testing.T) {
	arrs :=[]int{5,4,3,2,1}
	InsertionSort(arrs)
	expected:=[]int{1,2,3,4,5}
	if !reflect.DeepEqual(arrs,expected){
		t.Errorf("not equal")
	}
}
