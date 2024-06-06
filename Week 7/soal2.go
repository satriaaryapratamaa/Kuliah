package main
import (
	"fmt"
)
const NMax int = 100
type tabChar [NMax]byte


func main(){
	var n int
	var T tabChar
    
	masukanArray(&T, &n)
	lowerCase(&T, n)
	cetakArray(T, n)
}

func masukanArray(T *tabChar, n *int){
    /*I.S. Data tersedia pada piranti masukan
    Proses: Masukan akan terus berlangsung dan berhenti ketika pengguna
    memasukkan '.'
    F.S. Sekumpulan karakter sebanyak n berada dalam array T
    Petunjuk: Gunakan while-loop untuk melakukan proses input
    */
    var char byte
    
    fmt.Scanf("%c",&char)
    for char != '.'{
        (*T)[*n] = char
		*n++
        fmt.Scanf("%c", &char)
    }
}

func lowerCase(T *tabChar,n int){
    /*I.S. Terdefinisi sekumpulan n karakter pada array T
    F.S. Seluruh anggota karakter pada array T dikonversi menjadi huruf kecil
    Petunjuk: Gunakan ASCII, apabila sudah huruf kecil, tidak perlu diubah
    */
	for i := 0; i < n; i++ {
		if 'A' <= T[i] && T[i] <= 'Z' {
			T[i] = T[i] + 32
		}
	}
}

func cetakArray(T tabChar, n int){
    /*I.S. Terdefinisi sekumpulan n karakter pada array T
    F.S. Menampilkan seluruh elemen array T pada layar
    */
	for i := 0; i < n; i++ {
		fmt.Printf("%c", T[i])
	}
}