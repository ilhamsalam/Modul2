package main

import "fmt"

func main() {
        var bunga string
        var pita string
        var jumlahBunga int

        for {
                fmt.Printf("Bunga %d : ", jumlahBunga+1)
                fmt.Scan(&bunga)

                if bunga == "SELESAI" {
                        break
                }

                pita += bunga + " - "
                jumlahBunga++
        }

        if len(pita) > 0 {
                pita = pita[:len(pita)-3]
        }

        fmt.Println("Pita : ", pita)
        fmt.Println("Bunga : ", jumlahBunga)
}