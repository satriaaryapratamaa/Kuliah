package main
import "fmt"

const NMAX int = 2022

type nasabah struct {
	id, nama, bank string
	rek int
}

type tabNasabah [NMAX]nasabah

func addNasabah(p *tabNasabah, n int) {
	for i := 0; i < n; i++ {
		fmt.Scan(&p[i].id, &p[i].nama, &p[i].bank, &p[i].rek)
	}
}

func cetak(p nasabah, n int, k string) {
	if n < NMAX {
		for i := 0; i < n-1; i++ {
			
		}
	} else {
		fmt.Print("data penuh")
	}
}

func main() {
	var t tabNasabah
	var n int
	var x string

	fmt.Scan(&n)
	addNasabah(&t, n)
	cetak(t,n,x)
}
