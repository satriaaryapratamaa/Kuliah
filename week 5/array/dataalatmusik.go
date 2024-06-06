package main

import (
	"fmt"
)

// Definisikan tipe bentukan untuk alat musik
type AlatMusik struct {
	jenis string
	merk  string
	tipe  string
	harga int
}

func main() {
	// Slice untuk menampung data alat musik
	alatMusiks := make([]AlatMusik, 0)

	// Input data alat musik awal
	for i := 0; i < 3; i++ {
		var jenis, merk, tipe string
		var harga int
		fmt.Scanf("%s %s %s %d", &jenis, &merk, &tipe, &harga)
		alatMusiks = append(alatMusiks, AlatMusik{jenis, merk, tipe, harga})
	}

	// Input data alat musik tambahan sampai jenis "selesai"
	for {
		var jenis, merk, tipe string
		var harga int
		fmt.Scanf("%s", &jenis)
		if jenis == "selesai" {
			break
		}
		fmt.Scanf("%s %s %d", &merk, &tipe, &harga)
		alatMusiks = append(alatMusiks, AlatMusik{jenis, merk, tipe, harga})
	}

	// Map untuk menampung alat musik termurah berdasarkan jenis
	musikTermurah := make(map[string]AlatMusik)

	// Cari alat musik termurah untuk setiap jenis
	for _, alat := range alatMusiks {
		if termurah, exists := musikTermurah[alat.jenis]; !exists || alat.harga < termurah.harga {
			musikTermurah[alat.jenis] = alat
		}
	}

	// Tampilkan alat musik termurah untuk setiap jenis
	for _, jenis := range []string{"Gitar", "Drum", "Piano"} {
		if alat, exists := musikTermurah[jenis]; exists {
			fmt.Println(alat.jenis, alat.merk, alat.tipe, alat.harga)
		}
	}
}
