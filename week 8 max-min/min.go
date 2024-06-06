package main
import "fmt"

const NMAX = 256
type Arr[NMAX]int

func input(x *Arr, n *int) {
	fmt.Scan(n)
	for i := 0; i < *n; i++ {
		fmt.Scan(&x[i])
	}
}

func returnData(x Arr, n int) int {
/*
	var min int
	min = x[0]
	for i := 0; i < n; i++ {
		if min < x[i] {
			min = x[i]
		}
	}
	return min
*/
	
	// MENENTUKAN POSISI URUTAN KE
	var posisi int
	posisi = 0

	for i := 1; i < n; i++ {
		if x[posisi] < x[i] {
			posisi = i 
		}
	}
	return posisi // Karena array diitung dari 0... n, maka kita perlu tambah agar sesuai dengan bahasa manusia//
}

func main() {
	var x Arr
	var n int

	input(&x,&n)
	fmt.Println(returnData(x,n))
}