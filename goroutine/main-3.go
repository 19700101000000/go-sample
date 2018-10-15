package main

import "fmt"

type channels struct {
	startA chan string
	startB chan string
	finish chan string
}

func main() {
	chs := &channels{
		startA: make(chan string),
		startB: make(chan string),
		finish: make(chan string),
	}
	go b(chs)
	go a(chs)

	f := <-chs.finish

	fmt.Println(f)
}

func a(chs *channels) {
	fmt.Println("this is A.")

	chs.startB <- "foo"
	bar := <-chs.startA

	fmt.Println(bar)

	chs.startB <- ""
}

func b(chs *channels) {
	foo := <-chs.startB

	fmt.Println(foo)

	chs.startA <- "bar"
	<-chs.startB

	fmt.Println("this is B.")
	chs.finish <- "fin."
}
