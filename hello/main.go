package main

import "fmt"

func main() {
	str1, str2 := "hello", "world"
	//var str1, str2 string = "hello", "world"
	fmt.Println(str1, str2)

	str1, str2 = str2, str1
	fmt.Println(str1, str2)
}
