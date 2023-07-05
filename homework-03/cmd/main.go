package main

import (
	"fmt"
	"go-course-4/homework-03/pkg/list"
)

func main() {
	l := list.New()
	l.Push(list.Elem{Val: 2})
	l.Push(list.Elem{Val: 1})
	fmt.Println(l)

}
