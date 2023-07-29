package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	min := 1.0
	max := 100.0
	minChance := max * 0.2
	randNum := min + rand.Float64()*(max-min)

	ch := make(chan string)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		player("Player 1", ch, minChance, randNum)
		wg.Done()
	}()
	go func() {
		player("Player 2", ch, minChance, randNum)
		wg.Done()
	}()

	ch <- "start"
	wg.Wait()
}

func player(name string, table chan string, minChance float64, randNum float64) {
	status := <-table
	for {
		if status == "start" {
			fmt.Println("Someone lost", name)
			return
		}

		//table <- "stop"
		//if status == "stop" {
		//	fmt.Println("Someone lost", name)
		//	//return
		//}

		//close(table)
		//if randNum < minChance {
		//	fmt.Printf("Player %s missed\n", name)
		//	close(table)
		//	return
		//}
	}
}
