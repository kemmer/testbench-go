package main

import "fmt"

type Foo interface {
	TestMethod(v int) int
}

type MyEcho struct {
}

func (e *MyEcho) TestMethod(v int) int {
	return v + 3
}

func main() {
	me := MyEcho{}

	fmt.Println("aaa")
	fmt.Println("bbb", me.TestMethod(88))
}
