package main

import "fmt"

// Subprogram prosedur untuk menghitung total luas dan total aset
func hitungProperti(x int, totalLuas *int, totalAset *int) {
    var pj, lb, perkav int

    for i := 0; i < x; i++ {
        fmt.Printf("Masukkan panjang, lebar, dan harga perkavling untuk properti %d: ", i+1)
        fmt.Scan(&pj, &lb, &perkav)
        *totalLuas = pj * lb + *totalLuas
        *totalAset = *totalLuas * perkav + *totalAset
    }
}

func main() {
    var n, luas, aset int

    fmt.Print("Masukkan jumlah properti yang dimiliki: ")
    fmt.Scan(&n)

    hitungProperti(n, &luas, &aset)

    fmt.Printf("Total luas properti yang dimiliki: %d meter persegi\n", luas)
    fmt.Printf("Total aset yang dimiliki: %d\n", aset)
}
