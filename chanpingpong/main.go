package main

import (
	"fmt"
	"time"
)

func pinger(c chan string) {
	for i := 0; ; i++ {
		c <- "ping"
	}
}

func ponger(c chan string) {
	for i := 0; ; i++ {
		c <- "pong"
	}
}

func printer(c chan string) {
	i := 0
	for {
		i++
		msg := <-c
		fmt.Println(msg, i)
		//time.Sleep(time.Second * 1)
	}
}

func main() {
	var c chan string = make(chan string)

	go pinger(c)
	go ponger(c)
	go printer(c)
	time.Sleep(time.Millisecond * 1)

	//var input string
	//fmt.Scanln(&input)
}
