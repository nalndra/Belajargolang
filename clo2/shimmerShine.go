package main

import "fmt"

func main() {
    var y ,count0, countNum, num int

    fmt.Scan(&y)

    for i := 0; i < 9; i++ {
        fmt.Scan(&num)
        if num == 0 {
            count0++
        } else if num == y {
            countNum++
        }
    }

    if count0 > countNum {
        fmt.Println("Median bernilai 0")
    } else {
        fmt.Println("Median bernilai", y)
    }
}