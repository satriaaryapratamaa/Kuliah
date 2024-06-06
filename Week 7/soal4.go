package main

import "fmt"

type kubus struct {
    sisi, luasP, volume int
}
type arrKubus [100]kubus

func hitungKubus(k *arrKubus, n int) {
    var totV, totP int
    for i := 0; i < n; i++ {
        k[i].luasP = 6 * k[i].sisi * k[i].sisi
        k[i].volume = k[i].sisi * k[i].sisi * k[i].sisi
        totV += k[i].volume
        totP += k[i].luasP
    }
    fmt.Println(totP, totV)
}

func main() {
    var k arrKubus
    var n, sisi int
    fmt.Scan(&sisi)
    for sisi != 0 {
        k[n].sisi = sisi
        n++
        fmt.Scan(&sisi)
    }
    hitungKubus(&k, n)
}