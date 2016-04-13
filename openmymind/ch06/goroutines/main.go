package main

import (
	"fmt"
	"sync"
	"time"
)

var counter = 0
var lock sync.Mutex

func main() {
	/*fmt.Println("start")
	go process()
	time.Sleep((time.Millisecond * 5000))
	go func() {
		fmt.Println("p2")
	}()
	fmt.Println("done")*/

	/*for i := 0; i < 10; i++ {
		go incr(i)
	}
	time.Sleep(time.Millisecond * 2000)*/

	for i := 0; i < 10; i++ {
		go incr1(i)
	}
	time.Sleep(time.Millisecond * 2000)

}

func incr(i int) {
	counter++
	fmt.Println(i, counter)
}

func incr1(i int) {
	lock.Lock()
	counter++
	fmt.Println(i, counter)
	defer lock.Unlock()
}

func process() {
	fmt.Println("processing")
}
