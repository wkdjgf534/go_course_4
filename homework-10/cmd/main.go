package main

import (
	"fmt"
	"sync"

	"go-course-4/homework-10/pkg/player"
)

func main() {
	var wg sync.WaitGroup

	ch := make(chan string)
	p1 := player.New("Player 1")
	p2 := player.New("Player 2")

	wg.Add(2)
	go p1.Play(ch, &wg)
	go p2.Play(ch, &wg)

	ch <- "begin"
	wg.Wait()

	fmt.Printf("%s, score: %d\n", p1.Name, p1.Point)
	fmt.Printf("%s, score: %d\n", p2.Name, p2.Point)
}
