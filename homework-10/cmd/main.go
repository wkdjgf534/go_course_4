package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	ch := make(chan string)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		player("Player 1", ch)
		wg.Done()
	}()
	go func() {
		player("Player 2", ch)
		wg.Done()
	}()

	ch <- "begin"
	wg.Wait()
}

func player(name string, table chan string) {
	min := 1.0
	max := 100.0
	point := 1
	for t := range table {
		minChance := max * 0.2
		randNum := min + rand.Float64()*(max-min)

		if t == "begin" {
			fmt.Println("Player ", name, "lose!")
			table <- "ping"
		}

		fmt.Println(randNum, minChance)
		if randNum < minChance {
			fmt.Printf("Player %s missed\n", name)
			table <- "stop"
		}

		if t == "stop" {
			close(table)
			return
		}

		if t == "ping" {
			fmt.Println(t)
			table <- "pong"
		}

		if t == "pong" {
			fmt.Println(t)
			table <- "ping"
		}
	}
}
