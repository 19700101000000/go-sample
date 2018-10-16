package main

import "fmt"

func main() {
	ch := make(chan struct{})

	go foo(ch)

	if _, ok := <-ch; !ok {
		fmt.Println("channel is close")
		return
	}
}

func foo(ch chan struct{}) {
	fmt.Println("hello, goroutine.")
	close(ch)
}
