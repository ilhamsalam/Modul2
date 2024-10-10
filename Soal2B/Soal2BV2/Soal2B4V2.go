package main

import "fmt"

func f(k int) float64 {
    pembilang := float64((4*k + 2) * (4*k + 2))
    penyebut := float64((4*k + 1) * (4*k + 3))
    return pembilang / penyebut
}

func main() {
    var K int
    fmt.Print("Nilai K = ")
    fmt.Scan(&K)

    hasil := 1.0
    for k := 0; k < K; k++ {
        hasil *= f(k)
    }

    fmt.Printf("Nilai akar 2 = %.10f\n", hasil)
}
