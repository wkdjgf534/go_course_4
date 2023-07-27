package main

import (
	"fmt"
	"time"
)

func loop(msg string) {
	for {
		fmt.Println(msg)
		time.Sleep(time.Second)
	}
}

func main() {
	go loop("ping")
	loop("pong")
}
