package main

import "fmt"
import "strings"

func main() {

    s := "/users/1"
    slice := strings.Split(s, "/")
    fmt.Println(len(slice))
    for _, x := range slice {
        fmt.Println(x)
    }
}
