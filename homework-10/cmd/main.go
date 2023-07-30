package main

import (
	"fmt"
	"math/rand"
	"sync"
)

const MaxPoint = 10

func main() {
	ch := make(chan string)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		p := player("Player 1", ch)
		fmt.Println("Player 1", p)
		wg.Done()
	}()
	go func() {
		p := player("Player 2", ch)
		fmt.Println("Player 2", p)
		wg.Done()
	}()

	ch <- "begin"
	wg.Wait()
}

func player(name string, table chan string) int {
	min := 1.0
	max := 100.0
	point := 1
	for t := range table {
		minChance := max * 0.2
		randNum := min + rand.Float64()*(max-min)

		switch t {
		case "begin":
			table <- Kick(t)
		case "stop":
			fmt.Println(name, "lose!")
			table <- Kick(t)
		case "ping", "pong":
			fmt.Println(t)

			if randNum < minChance {
				point++
				if point >= MaxPoint {
					close(table)
				} else {
					table <- "stop"
				}

			} else {
				table <- Kick(t)
			}
		}
	}
	return point
}

func Kick(status string) string {
	if status == "ping" {
		return "pong"
	}
	return "ping"
}
