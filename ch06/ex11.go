package main

import "fmt"

func zero(x *int) {
    *x = 0
}

func one(x *int) {
    *x = 1
}

func main() {
    x := 5
    zero(&x)
    fmt.Println(x)

    xPtr := new(int)
    one(xPtr)
    fmt.Println(*xPtr)
}
