package main

import (
	"fmt"
	"time"
)

const maxScore = 10

func loop(msg string) {
	for {
		fmt.Println(msg)
		time.Sleep(time.Second)
	}
}

func main() {
	//ch := make(chan string)
	//ch <- "begin"
	//ch <- "stop"
	//var wg sync.WaitGroup

	//wg.Add()
	go loop("ping")
	loop("pong")
	//wg.Wait()
}
