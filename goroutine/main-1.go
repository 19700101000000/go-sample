package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg *sync.WaitGroup = new(sync.WaitGroup)
	wg.Add(1)
	ch := make(chan bool)
	go a(ch, wg)
	go b(ch, wg)
	wg.Wait()
}

func a(ch chan bool, wg *sync.WaitGroup) {
	time.Sleep(1 * time.Second)
	ch <- true
}

func b(ch chan bool, wg *sync.WaitGroup) {
	<-ch
	fmt.Println("hello, world.")
	wg.Done()
}
