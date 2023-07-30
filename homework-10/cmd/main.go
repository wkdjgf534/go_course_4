package main

import (
	"fmt"
	"math/rand"
	"sync"

	"go-course-4/homework-10/pkg/player"
)

const MaxPoint = 10

func main() {
	ch := make(chan string)
	p1 := player.New("Player 1")
	p2 := player.New("Player 2")

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		play(p1, ch)
		wg.Done()
	}()
	go func() {
		play(p2, ch)
		wg.Done()
	}()

	ch <- "begin"
	wg.Wait()
	fmt.Printf("%s, score: %d\n", p1.Name, p1.Point)
	fmt.Printf("%s, score: %d\n", p2.Name, p2.Point)
}

func play(player *player.Player, table chan string) {
	min := 1.0
	max := 100.0
	for t := range table {
		minChance := max * 0.2
		randNum := min + rand.Float64()*(max-min)

		switch t {
		case "begin", "stop":
			table <- Kick(t)
		case "ping", "pong":
			fmt.Println(t)

			if randNum >= max-minChance { //successfull strike
				player.Point++
				if player.Point >= MaxPoint {
					fmt.Printf("player: %s won\n", player.Name)
					close(table)
				} else {
					table <- "stop"
				}

			} else {
				table <- Kick(t)
			}
		}
	}
}

func Kick(status string) string {
	if status == "ping" {
		return "pong"
	}
	return "ping"
}
