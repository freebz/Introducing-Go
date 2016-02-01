package main

import "fmt"

func main() {
	var x string = "hello"
	var y string = "world"
	var z string = "hello"
	fmt.Println(x == y)
	fmt.Println(x == z)
}
