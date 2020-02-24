package main

import (
	"fmt"
	"reflect"
)
func main() {
	fmt.Println(reflect.TypeOf((*error)(nil)).Elem())
	fmt.Println(reflect.TypeOf((error)(nil)))
}