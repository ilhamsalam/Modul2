package main

import "fmt"

func kabisat(tahun int) bool {
        return (tahun%400 == 0) || (tahun%4 == 0 && tahun%100 != 0)
}

func main() {
        var tahun int

        fmt.Print("Tahun : ")
        fmt.Scanln(&tahun)

        fmt.Println(kabisat(tahun))
}