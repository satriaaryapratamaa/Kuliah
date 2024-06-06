// Buatlah prosedur dalam bahasa Go untuk menghitung banyaknya huruf selain huruf 'a' dan 'i' pada rangkaian karakter. 

// procedure sum(kar byte, jumlah int) {

// I.S. Terdefinisi kar byte dan jumlah int

// F.S. jumlah huruf selain  'a' dan 'i' terdefinisi di jumlah int 

// }

// Masukan berupa rangkaian karakter dengan karakter titik sebagai sentinel.

// Keluaran berupa jumlah huruf selain  'a' dan 'i' yang terdapat pada rangkaian karakter itu.

package main
import "fmt"

func sum(kar string, total int) {
	
	for i := 0; i < len(kar); i++ {
        if kar [i:i+1] != "a" && kar[i:i+1] != "i" && kar[i:i+1] != "." {
            total++
        } else {
			fmt.Print("")
		}
    }
	fmt.Println(total)
}

func main() {
	var x string
	var total int

	fmt.Scan(&x)
	sum(x, total)
}

// VERSI 2 //

/*
package main

import "fmt"

func sum(kar byte, jumlah *int) {
	for kar != '.' {
		if kar != 'a' && kar != 'i' {
			*jumlah++
		}
		fmt.Scanf("%c", &kar)
	}
}

func main() {
	var jumlah int
	var kar byte
	fmt.Scanf("%c", &kar)
	sum(kar, &jumlah)
	fmt.Println(jumlah)
}
*/