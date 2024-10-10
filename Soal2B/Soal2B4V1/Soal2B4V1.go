package main

import "fmt"

func f(k int) float64 {
    pembilang := float64((4*k + 2) * (4*k + 2))
    penyebut := float64((4*k + 1) * (4*k + 3))
    return pembilang / penyebut
}

func main() {
    var k int
    fmt.Print("Nilai K = ")
    fmt.Scan(&k)

    hasil := f(k)
    fmt.Printf("Nilai f(K) = %.10f\n", hasil)
}
