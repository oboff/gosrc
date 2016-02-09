// main
package main

import (
	"fmt"
	"log"
	"runtime"
	"time"
)

func main() {
	where := func() {
		_, file, line, _ := runtime.Caller(1)
		log.Printf("%s:%d", file, line)
	}
	x := 15
	where()
	fmt.Println(x)
	log.SetFlags(log.Llongfile)
	log.Print("")
	time.Now().Sub()
}
