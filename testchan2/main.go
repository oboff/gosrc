package main

import (
	"fmt"
	"time"
)

func simulateEvent(name string, timeInSecs time.Duration, timeNow *time.Time) {
	// sleep for a while to simulate time consumed by event
	fmt.Println("Started ", name, ": Should take", timeInSecs*time.Second, "seconds.")
	timeElapsed := time.Since(*timeNow)
	time.Sleep(timeInSecs)
	fmt.Println("Finished ", name, "Time elapsed: ", timeElapsed)
}

func main() {
	now := time.Now()
	go simulateEvent("100m sprint", 10, &now) //start 100m sprint, it should take 10 seconds
	go simulateEvent("Long jump", 6, &now)    //start long jump, it should take 6 seconds
	go simulateEvent("High jump", 3, &now)    //start high jump, it should take 3 seconds
	fmt.Printf("Total time: %v\n", time.Since(now))
	//so that the program doesn't exit here, we make the program wait here for a while
	time.Sleep(1200 * time.Millisecond)
}
