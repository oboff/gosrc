// main
package main

import (
	"fmt"
)

func main() {
	var num, den int
	fmt.Print("Enter numerator: ")
	fmt.Scanf("%d", &num)
	fmt.Print("Enter denominator: ")
	fmt.Scanf("%d", &den)
	fmt.Printf("%d / %d = %d %d/%d (= %f)\n", num, den, num/den, num%den, den, float64(num)/float64(den))
}
