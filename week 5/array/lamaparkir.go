//definisikan tipe bentukan struktur untuk menyatakan data waktu dengan komponen jam, menit,
//dan detik. setelah itu buatlah program untuk menghitung lama waktu kendaraan parkir.

package main
import "fmt"

type lama_waktu struct {
	jam, menit, detik int
}

func main() {
	var t2, t1, t lama_waktu 
	var total int
	fmt.Scan(&t2.jam, &t2.menit, &t2.detik)
	fmt.Scan(&t1.jam, &t1.menit, &t1.detik)
	
	total = (t2.jam - t1.jam)*3600 + (t2.menit - t1.menit)*60 + (t2.detik - t1.detik)
	
	t.jam = total / 3600
	t.menit = total % 3600 / 60 
	t.detik = total % 3600 % 60 
	
	if t.jam < 0 && t.menit < 0 && t.detik < 0 {
		t.jam = total / 3600 * -1
		t.menit = total % 3600 / 60 * -1 
		t.detik = total % 3600 % 60 * -1
	}
	
	fmt.Println(t.jam, t.menit, t.detik)
}