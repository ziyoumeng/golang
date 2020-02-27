package main

import (
	"fmt"
	"github.com/ziyoumeng/golang/arithmetic/LRU"
	"github.com/ziyoumeng/golang/arithmetic/linked_list"
)

func main() {
	doublyList := linked_list.NewDoublyLinkedList()
	doublyList.Prepend(linked_list.NewNode(1, 1))
	doublyList.Prepend(linked_list.NewNode(2, 2))
	doublyList.Print()
	doublyList.Delete(doublyList.Search(1))
	doublyList.Print()

	cache := LRU.NewLRU(3)
	cache.Put(1, 1)
	cache.Put(2, 2)
	cache.Put(3, 3)
	cache.Put(4, 4)
	cache.Print()
	testGet(cache, 2)
	testGet(cache, 4)
}

func testGet(cache *LRU.LRU, key interface{}) {
	val, ok := cache.Get(key)
	fmt.Println(key, val, ok)
	cache.Print()
}
