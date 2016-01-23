// main
package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)
	m := make(map[string]string)

	go func() {
		for {
			c <- "ping"
			time.Sleep(2 * time.Millisecond)
		}
	}()
	go func() {
		for {
			c <- "pong"
			time.Sleep(2 * time.Millisecond)
		}
	}()
	go func() {
		for {
			s := <-c
			switch s {
			case "ping":
				m[s] = m[s] + "x"
			case "pong":
				m[s] = m[s] + "o"
			}
			fmt.Printf("%s:%s\n", s, m[s])

		}
	}()

	time.Sleep(10 * time.Millisecond)
}
