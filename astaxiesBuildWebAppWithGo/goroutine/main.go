package main

import (
	"fmt"
	"runtime"
)

func say(s string) {
	for i := 0; i < 10; i++ {
		runtime.Gosched()
		fmt.Println(s)
	}
}

func sum(a []int, c chan int) {
	total := 0
	for _, v := range a {
		total += v
	}

	c <- total
}

func fibonacic(n int, c chan int) {

	x, y := 1, 1
	for i := 0; i < n; i++ {
		// fmt.Println("X = ", x)
		c <- x
		x, y = y, x+y
	}

	close(c)
}

func fibonacic1(c, quit chan int) {
	x, y := 1, 1
	for {
		select {
		case c <- x:
			fmt.Println("x = ", x)
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	//go say("world")
	// say("hello")

	a := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)

	x, y := <-c, <-c

	fmt.Println(x, y, x+y)

	c1 := make(chan int, 10)
	go fibonacic(cap(c1), c1)
	for i := range c1 {
		fmt.Println(i)
	}

	c2 := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("i = ", <-c2)
		}
		quit <- 0
	}()
	fibonacic1(c2, quit)
}
