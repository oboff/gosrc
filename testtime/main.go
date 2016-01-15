// main
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello World!")
	now := time.Now()
	fmt.Printf("%v\n", now)
	fmt.Printf("%v\n", time.Since(now))
}
