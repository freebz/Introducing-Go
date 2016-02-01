package main

import "fmt"

func main() {
    x := make(map[string]int)
    x["key"] = 10
    fmt.Println(x)
    fmt.Println(x["key"])

    y := make(map[int]int)
    y[1] = 10
    fmt.Println(y[1])
}
