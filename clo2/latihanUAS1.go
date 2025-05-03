package main

import (
    "fmt"
)

func main() {
    var number, original, reversed int
    fmt.Scan(&number)

    original = number

    for number > 0 {
        digit := number % 10
        reversed = reversed*10 + digit
        number /= 10
    }

    if original == reversed {
        fmt.Println("YA")
    } else {
        fmt.Println("TIDAK")
    }
}