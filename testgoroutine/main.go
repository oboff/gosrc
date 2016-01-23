package main

import (
	"fmt"
	"time"
)

func simulateEvent(name string, timeInSecs time.Duration) {
	// sleep for a while to simulate time consumed by event
	fmt.Println("Started ", name, ": Should take", timeInSecs*time.Second, "seconds.")
	now := time.Now()
	time.Sleep(timeInSecs * time.Second)
	fmt.Println("Finished ", name, "Time elapsed: ", time.Since(now).Seconds())
}

func main() {
	go simulateEvent("100m sprint", 10) //start 100m sprint, it should take 10 seconds
	go simulateEvent("Long jump", 6)    //start long jump, it should take 6 seconds
	go simulateEvent("High jump", 3)    //start high jump, it should take 3 seconds
	//so that the program doesn't exit here, we make the program wait here for a while
	//time.Sleep(12 * time.Second)
	for c := 0; c < 10+1; c++ {
		time.Sleep(980 * time.Millisecond)
		fmt.Printf("%ds...", c)
	}
}
