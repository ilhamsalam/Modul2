package main

import "fmt"

func main() {
    var n int
    var bunga string
    var pita string

    fmt.Print("N : ")
    fmt.Scan(&n)

    if n <= 0 {
        fmt.Println("N harus bilangan bulat positif")
        return
    }

    for i := 1; i <= n; i++ {
        fmt.Printf("Bunga %d : ", i)
        fmt.Scan(&bunga)
        pita += bunga + " - "
    }

    if len(pita) > 0 {
        pita = pita[:len(pita)-3]
    }

    fmt.Println("Pita : ", pita)
}
