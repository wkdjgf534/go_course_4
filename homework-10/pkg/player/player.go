package player

import (
	"fmt"
	"math/rand"
	"sync"
)

const maxPoint = 10

type Player struct {
	Name  string
	Point int
}

// New - create a new player
func New(name string) *Player {
	return &Player{
		Name:  name,
		Point: 0,
	}
}

// Play - generate a game's process
func (p *Player) Play(table chan string, wg *sync.WaitGroup) {
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
