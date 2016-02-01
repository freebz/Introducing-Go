package main

import "fmt"

func main() {
    elements := map[string] map[string]string{
        "H": map[string]string{
            "name": "Hydrogen",
            "state", "gas",
        },
        "He": map[string]string{
            "name": "Helium",
            "state": "gas",
        },
    }

    fmt.Println(elements["Li"])
    fmt.Println(elements["Un"])
    
    el, ok := elements["Un"]
    el["name"] = "Un t"

    fmt.Println(el["name"], ok)

    el, ok = elements["Un"]
    fmt.Println(el["name"], ok)
}
