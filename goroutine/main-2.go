package main

import (
	"fmt"
)

type foo struct {
	x string
	y string
}

func main() {
	ch := make(chan foo)

	go z(ch)

	f := <-ch
	fmt.Printf("%s %s\n", f.y, f.x)
}

func z(ch chan foo) {
	ch <- foo{
		x: "hello",
		y: "world",
	}
}
