package main

import "fmt"

func main() {
    i := 1
    for i <= 3 {
        fmt.Println(i)
        i = i + 1
    }

    for n := 0; n <= 5; n++ {
        fmt.Println(n)
    }
}
