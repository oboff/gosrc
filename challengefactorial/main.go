// main
package main

import (
	"fmt"
)

func fact(n int) int {
	total := 1
	for i := n; i > 0; i-- {
		total *= i
		//fmt.Println(i, x)
	}
	return total
}

func factrec(n int) int {
	if n > 1 {
		return factrec(n-1) * n
	} else {
		return n
	}
}

func xfactValidator(b, a int) bool {
	if a > b || b < 0 {
		fmt.Errorf("Error: %d > %d", a, b)
		return false
	}
	return true
}

func factgo(n int) chan int {
	c := make(chan int)
	go func() {
		total := 1
		for i := n; i > 0; i-- {
			total *= i

		}
		c <- total
		close(c)
	}()
	return c
}

func factr(b int, a int, c chan int) chan int {
	//fmt.Println(b, a)
	total := 1
	func() {
		for i := b; i > a; i-- {
			//fmt.Printf("i:%d\n", i)
			if i < 2 {
				break
			}
			total *= i
			//fmt.Println(total)
		}
		c <- total
	}()
	return c
}

func decrementor(n, m int) chan int {
	out := make(chan int)
	go func() {
		for i := n; i > m; i-- {
			out <- i
		}
		close(out)
	}()
	return out
}

func multiplier(c chan int) chan int {
	out := make(chan int)
	go func() {
		sum := 1
		for n := range c {
			sum *= n
		}
		out <- sum
		close(out)
	}()
	return out
}

func main() {
	num := 12
	chans := 4
	step := num / chans

	ch := make(chan chan int, chans)

	total := 1

	go func() {
		start := num
		for i := chans; i > 0; i-- {
			ch <- multiplier(decrementor(start, start-step))
			start -= step
		}
		close(ch)
	}()

	for i := range ch {
		for j := range i {
			total *= j
		}
	}

	fmt.Println(total)
	fmt.Println(fact(num))
	//	c := make(chan chan int, 2)
	//	c <- multiplier(decrementor(5, 3))
	//	c <- multiplier(decrementor(3, 0))
	//	close(c)
	//	total := 1
	//	for n := range c {
	//		for m := range n {
	//			total *= m
	//		}
	//	}
	//	fmt.Println(total)

	//	fmt.Println("Hello World!")
	//	fmt.Println(fact(4))
	//	fmt.Println(factrec(4))
	//	fmt.Println(factr(4, 3))
	//	fmt.Println(factr(3, 2))
	//	fmt.Println(factr(2, 1))
	//	fmt.Println(factr(1, 0))

	//	c := make(chan int)

	//	go factc(2, 0, c)
	//	go factc(4, 2, c)
	//	for n := range c {
	//		fmt.Println(n)
	//	}

	//	f := factgo(5)
	//	for n := range f {
	//		fmt.Println(n)
	//	}

	//	c := make(chan int, 2)
	//	c <- 1
	//	c <- 2
	//	close(c)
	//	for n := range c {
	//		fmt.Println(n)
	//	}

	//	n := 10
	//	s := 4
	//	ch := make(chan int, s)
	//	func() {
	//		x := n
	//		fmt.Printf("x:%d\n", x)
	//		s1 := s
	//		fmt.Printf("s1:%d\n", s1)
	//		for i := s1; i > 0; i-- {

	//			f := factr(x, x-(x/s1))
	//			fmt.Printf("x:%d, x-x/s1:%d, x*x-(x/s1):%d\n", x, x-(x/s1), f)
	//			ch <- f
	//			x -= x / s1
	//		}
	//	}()
	//	close(ch)

	//	total := 1
	//	for n := range ch {
	//		total *= n
	//	}
	//	fmt.Println("total:", total)

	//fmt.Printf("end:%d\n", factr(4, 2))
}
