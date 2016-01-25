package main

import (
	"fmt"
	"sync"
)

func fact(n int) int {
	if n > 0 {
		return fact(n-1) * n
	}
	return 1
}

func gen(nums ...int) chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func factc(in chan int) chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- fact(n)
		}
		close(out)
	}()
	return out
}

func merge(cs ...chan int) chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	output := func(c chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}

	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func main() {
	nums := []int{}
	for i := 1; i < 1001; i++ {
		nums = append(nums, i%13)
	}

	g := gen(nums...)
	chans := []chan int{}
	for i := 1; i < 8; i++ {
		chans = append(chans, factc(g))
	}

	index := 1
	for n := range merge(chans...) {
		fmt.Println(index, ":", n)
		index++
	}
}
