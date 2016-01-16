// main
package main

import (
	"fmt"
)

func sum35(n int) int {
	sum := 0
	for i := 0; i < n; i++ {
		switch {
		case i%3 == 0:
			sum += i
			//fmt.Println(i, i%3, sum)
		case i%5 == 0:
			sum += i
			//fmt.Println(i, i%5, sum)
		}
	}
	return sum

}

func main() {
	fmt.Println(sum35(1000))
}
