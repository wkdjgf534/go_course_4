package main

import (
	"fmt"
	"math/rand"
	"sync"

	"go-course-4/homework-10/pkg/player"
)

const maxPoint = 10

var wg sync.WaitGroup

func main() {
	ch := make(chan string)
	p1 := player.New("Player 1")
	p2 := player.New("Player 2")

	wg.Add(2)
	go play(p1, ch)
	go play(p2, ch)

	ch <- "begin"
	wg.Wait()

	fmt.Printf("%s, score: %d\n", p1.Name, p1.Point)
	fmt.Printf("%s, score: %d\n", p2.Name, p2.Point)
}

func play(p *player.Player, table chan string) {
	defer wg.Done()

	min := 1.0
	max := 100.0
	minChance := max * 0.2
	for t := range table {
		randNum := min + rand.Float64()*(max-min)

		switch t {
		case "begin", "stop":
			table <- kickBall(t)
		case "ping", "pong":
			fmt.Println(t)

			if randNum >= max-minChance {
				p.Point++
				if p.Point >= maxPoint {
					fmt.Printf("\nGame status: %s won\n", p.Name)
					close(table)
				} else {
					table <- "stop"
				}

			} else {
				table <- kickBall(t)
			}
		}
	}
}

func kickBall(status string) string {
	if status == "ping" {
		return "pong"
	}
	return "ping"
}
