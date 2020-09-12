package main

import (
 "fmt"
 "time"
)

func main() {
    var t = time.Now().Weekday()
    fmt.Println("I only drink Beer on Weekends")

    switch t {
        case time.Saturday, time.Sunday:
            fmt.Println("I drink Beer today")
        default:
            fmt.Println("I also drink Beer today")
    }
}
