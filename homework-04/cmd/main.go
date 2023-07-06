package main

import (
	"fmt"
	"go-course-4/homework-04/pkg/list"
)

func main() {
	l := list.New()
	l.Push(list.Elem{Val: 4})
	l.Push(list.Elem{Val: 3})
	l.Push(list.Elem{Val: 2})
	l.Push(list.Elem{Val: 1})
	fmt.Println(l)
	l.Pop()
	l.Pop()
	fmt.Println(l)

}
