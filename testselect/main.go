package main

import (
	"fmt"
	//"time"
)

var n int

func count() int {
	n++
	return n
}

func main() {
	c1 := make(chan string)
	c2 := make(chan string)
	c3 := make(chan int)

	go func() {
		for {
			c1 <- "from 1"
			c3 <- count()
			//time.Sleep(time.Second * 2)
		}
	}()

	go func() {
		for {
			c2 <- "from 2"
			c3 <- count()
			//time.Sleep(time.Second * 3)
		}
	}()

	go func() {
		for {
			select {
			case msg1 := <-c1:
				fmt.Println(msg1, <-c3)
			case msg2 := <-c2:
				fmt.Println(msg2, <-c3)
			}
		}
	}()
	//time.Sleep(time.Millisecond * 10)
	var input string
	fmt.Scanln(&input)
}
