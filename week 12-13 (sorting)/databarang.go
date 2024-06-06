package main

import "fmt"

const NMAX int = 250

type barang struct {
	nama, id      string
	jumlah, harga int
}

type tabBarang [NMAX]barang

func inputdata(d *tabBarang, n *int) {
	// akan menginput array barang, berhenti jika inputan none none 0 0
	stop := false
	var i int
	i = 0
	for !stop {
		fmt.Scan(&d[i].id, &d[i].nama, &d[i].jumlah, &d[i].harga)

		if d[i].id == "none" && d[i].nama == "none" && d[i].jumlah == 0 && d[i].harga == 0 {
			stop = true
		} else {
			i++
			*n++
		}
	}
}

func cetak(d tabBarang, n int) {
	//fmt.Println("%-s %-20s %-5s %-5s\n", "id", "nama", "jumlah", "harga")
	for i := 0; i < n; i++ {
		fmt.Printf("%-s %-20s %-5d %-5d\n", d[i].id, d[i].nama, d[i].jumlah, d[i].harga)
	}
	fmt.Println("")
}

func cariDataSeq(d tabBarang, n int, x string) {
	//mencari data berdasarkan id barangnya
	var ketemu bool
	ketemu = false

	for i := 0; i < n; i++ {
		if d[i].id == x {
			fmt.Println(d[i].nama)
			ketemu = true
			fmt.Println("")
		} 
	}
	if !ketemu {
		fmt.Println("data tidak ditemukan")
		fmt.Println("")
	}
}

func SelectionSort(d *tabBarang, n int) {
	//mengurutkan berdasarkan harga
	var pass, j, idx int
	var temp barang

	for pass = 0; pass < n-1; pass++{
		idx = pass
		for j = pass + 1; j < n; j++ {
			if d[j].harga < d[idx].harga {
				idx = j
			}
		}
		temp = d[pass]
		d[pass] = d[idx]
		d[idx] = temp
	}
	fmt.Println("")
}

func cariDataBin(d tabBarang, n int, h int) {
	// mencari data berdasarkan harga barangnya

	var bawah, atas, tengah int
	var ketemu bool
	bawah = 0
	atas = n-1
	ketemu = false

	for bawah <= atas && !ketemu {
		tengah = (bawah + atas)/ 2
		if h > d[tengah].harga {
			bawah = tengah + 1
		} else if h < d[tengah].harga {
			atas = tengah -1
		} else {
			ketemu = true
		}
	}
	fmt.Println(d[tengah].nama)
}

func InsertionSort(d *tabBarang, n int) {
	var pass, j int
	var temp barang
	for pass = 1; pass <= n; pass++ {
		j = pass
		temp = (*d)[pass]
		for j > 0 && (*d)[j-1] > temp {
			(*d)[j] = (*d)[j-1]
			j = j -1
		}
		(*d)[j] = temp
	}
}

func main() {
	var data tabBarang
	var x string
	var nData int
	var harga int

	inputdata(&data, &nData)
	cetak(data, nData)

	fmt.Scan(&x)
	cariDataSeq(data, nData, x)

	fmt.Scan(&harga)
	cariDataBin(data, nData, harga)

	SelectionSort(&data, nData)
	cetak(data,nData)

	InsertionSort(&data, nData)
	cetak(data,nData)
}
