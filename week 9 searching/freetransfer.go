package main

import "fmt"

type pemain struct {
	nama, posisi string
	umur         int
}

const NMAX int = 1040

type tabPemain [NMAX]pemain

func bacaPemain(Tp *tabPemain, nP, u int, p string) {
	for i := 0; i < nP; i++ {
		fmt.Scan(&Tp[i].nama, &Tp[i].umur, &Tp[i].posisi)
	}
}

func cariPemain(Tp *tabPemain, nP, u int, p string) int {
	var i, idx int
	idx = -1

	for i < nP && idx == -1 {
		if Tp[i].umur == u && Tp[i].posisi == p {
			idx = i
		} else {
			i++
		}
	}
	return idx
}

func main() {
	var Tp tabPemain
	var nP, u int
	var p string

	fmt.Print("Input Data Pemain : ")
	fmt.Scan(&nP)
	bacaPemain(&Tp, nP, u, p)

	fmt.Print("Input Data yang akan dicari, jika data ditemukan akan teroutput 1, dan sebaliknya -1 ")
	fmt.Scan(&u, &p)
	fmt.Println(cariPemain(&Tp, nP, u, p))
}
