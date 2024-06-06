package main

import "fmt"

type arr [30]int

func bacaselectionSort(t *arr, n int) {
	//fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&t[i])
	}
}

func selectionSort(t *arr, n int) {
	var pass, j, idx, temp int

	for pass = 0; pass < n-1; pass++ { // Loop dari pass ke 0 hingga n-2
		idx = pass
		for j = pass + 1; j < n; j++ { // Mencari elemen terkecil di sisa array
			if t[idx] > t[j] {
				idx = j
			}
		}
		// Menukar elemen terkecil dengan elemen pada posisi `pass`
		temp = t[idx]
		t[idx] = t[pass]
		t[pass] = temp
	}
}

func tampilArr(t arr, n int) {
	fmt.Println("Isi array setelah diurutkan:")
	for i := 0; i < n; i++ {
		fmt.Print(t[i], " ")
	}
	fmt.Println()
}

func main() {
	var t arr
	var n int

	fmt.Scan(&n)
	bacaselectionSort(&t, n)
	selectionSort(&t, n)
	tampilArr(t, n)

}
