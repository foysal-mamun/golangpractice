package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	c := make(chan int)
	for i := 0; i < 5; i++ {
		worker := &Worker{id: i}
		go worker.process(c)
	}

	for /*i := 0; i < 5; i++ */ {
		select {
		case c <- rand.Int():
		case <-time.After(time.Millisecond * 200):
			fmt.Println("time out")
		default:
			fmt.Println("dropped")

		}

		time.Sleep(time.Millisecond * 20)
	}
}

type Worker struct {
	id int
}

func (w *Worker) process(c chan int) {

	for {
		select {
		case data := <-c:
			fmt.Printf("worker %d got %d\n", w.id, data)
		case <-time.After(time.Millisecond * 10):
			fmt.Println("Break time")
			time.Sleep(time.Second)
		}

	}
}
