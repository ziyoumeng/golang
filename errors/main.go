package main

import (
	"github.com/pkg/errors"
	"fmt"
	"runtime"
)

func main() {
	err := middle()
	fmt.Printf("%+v", err)
}
func middle()error{
	err := source()
	return errors.Wrap(err,"middle")
}
func source()error{
	pc, file, line, ok := runtime.Caller(0)
	fmt.Println(pc,file,line,ok)

	fun := runtime.FuncForPC(pc+1)
	fmt.Println(fun.Name())
	panic(1)
	return errors.New("source")
}


