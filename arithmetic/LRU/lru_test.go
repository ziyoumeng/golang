package LRU

import (
	"reflect"
	"sync"
	"testing"
)

func TestLRU_Get(t *testing.T) {
	type fields struct {
		Head    *Node
		Tail    *Node
		Count   int
		MaxNum  int
		m       map[interface{}]*Node
		RWMutex sync.RWMutex
	}
	type args struct {
		key interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   interface{}
		want1  bool
	}{
		{
			name:   "empty lru",
			fields: fields{MaxNum: 3, m: make(map[interface{}]*Node)},
			args:   args{key: 1},
			want:   nil,
			want1:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LRU{
				Head:    tt.fields.Head,
				Tail:    tt.fields.Tail,
				Count:   tt.fields.Count,
				MaxNum:  tt.fields.MaxNum,
				m:       tt.fields.m,
				RWMutex: tt.fields.RWMutex,
			}
			got, got1 := l.Get(tt.args.key)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Get() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

//func TestLRU_Print(t *testing.T) {
//	type fields struct {
//		Head    *Node
//		Tail    *Node
//		Count   int
//		MaxNum  int
//		m       map[interface{}]*Node
//		RWMutex sync.RWMutex
//	}
//	tests := []struct {
//		name   string
//		fields fields
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			l := &LRU{
//				Head:    tt.fields.Head,
//				Tail:    tt.fields.Tail,
//				Count:   tt.fields.Count,
//				MaxNum:  tt.fields.MaxNum,
//				m:       tt.fields.m,
//				RWMutex: tt.fields.RWMutex,
//			}
//		})
//	}
//}
//
func TestLRU_PutAndGet(t *testing.T) {
	lru := NewLRU(3)
	lru.Put(1, 1)
	got, got1 := lru.Get(1)
	if !reflect.DeepEqual(got, 1) {
		t.Errorf("Get() got = %v, want %v", got, 1)
	}
	if got1 != true {
		t.Errorf("Get() got1 = %v, want %v", got1, true)
	}
	//type fields struct {
	//	Head    *Node
	//	Tail    *Node
	//	Count   int
	//	MaxNum  int
	//	m       map[interface{}]*Node
	//	RWMutex sync.RWMutex
	//}
	//type args struct {
	//	key   interface{}
	//	value interface{}
	//}
	//tests := []struct {
	//	name   string
	//	fields fields
	//	args   args
	//}{
	//	// TODO: Add test cases.
	//}
	//for _, tt := range tests {
	//	t.Run(tt.name, func(t *testing.T) {
	//		l := &LRU{
	//			Head:    tt.fields.Head,
	//			Tail:    tt.fields.Tail,
	//			Count:   tt.fields.Count,
	//			MaxNum:  tt.fields.MaxNum,
	//			m:       tt.fields.m,
	//			RWMutex: tt.fields.RWMutex,
	//		}
	//	})
	//}
}

//func TestLRU_deleteList(t *testing.T) {
//	type fields struct {
//		Head    *Node
//		Tail    *Node
//		Count   int
//		MaxNum  int
//		m       map[interface{}]*Node
//		RWMutex sync.RWMutex
//	}
//	type args struct {
//		node *Node
//	}
//	tests := []struct {
//		name   string
//		fields fields
//		args   args
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			l := &LRU{
//				Head:    tt.fields.Head,
//				Tail:    tt.fields.Tail,
//				Count:   tt.fields.Count,
//				MaxNum:  tt.fields.MaxNum,
//				m:       tt.fields.m,
//				RWMutex: tt.fields.RWMutex,
//			}
//		})
//	}
//}
//
//func TestLRU_moveToHead(t *testing.T) {
//	type fields struct {
//		Head    *Node
//		Tail    *Node
//		Count   int
//		MaxNum  int
//		m       map[interface{}]*Node
//		RWMutex sync.RWMutex
//	}
//	type args struct {
//		node *Node
//	}
//	tests := []struct {
//		name   string
//		fields fields
//		args   args
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			l := &LRU{
//				Head:    tt.fields.Head,
//				Tail:    tt.fields.Tail,
//				Count:   tt.fields.Count,
//				MaxNum:  tt.fields.MaxNum,
//				m:       tt.fields.m,
//				RWMutex: tt.fields.RWMutex,
//			}
//		})
//	}
//}
//
//func TestLRU_prependList(t *testing.T) {
//	type fields struct {
//		Head    *Node
//		Tail    *Node
//		Count   int
//		MaxNum  int
//		m       map[interface{}]*Node
//		RWMutex sync.RWMutex
//	}
//	type args struct {
//		node *Node
//	}
//	tests := []struct {
//		name   string
//		fields fields
//		args   args
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			l := &LRU{
//				Head:    tt.fields.Head,
//				Tail:    tt.fields.Tail,
//				Count:   tt.fields.Count,
//				MaxNum:  tt.fields.MaxNum,
//				m:       tt.fields.m,
//				RWMutex: tt.fields.RWMutex,
//			}
//		})
//	}
//}
//
//func TestNewLRU(t *testing.T) {
//	type args struct {
//		maxNum int
//	}
//	tests := []struct {
//		name string
//		args args
//		want LRU
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := NewLRU(tt.args.maxNum); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("NewLRU() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
