package main
import "fmt"

const NMAX int = 1024

type kelahiranBayi struct {
	wilayah string
	jumLahir, beratRendah, giziKurang int
}

type tabBayi [NMAX]kelahiranBayi

func bacaData(b *tabBayi, n int) {
	for i := 0; i < n; i++ {
		fmt.Scan(&b[i].wilayah, &b[i].jumLahir, &b[i].beratRendah, &b[i].giziKurang)
	}
}

func urutMembesar(b *tabBayi, n int) {
	var pass, idx, temp int

	for pass = 0; pass < n-1; pass++ {
		idx = pass
		for j := pass + 1; j < n - 1; j++ {
			if b[idx].jumLahir > b[j].jumLahir {
				idx = j
			}
		}
		temp = b[idx].jumLahir
		b[idx].jumLahir = b[pass].jumLahir
		b[pass].jumLahir = temp
	}
}

func tampilArrayBayi(b tabBayi, n int) {
	fmt.Println("urutan berdasarkan angka kelahiran kecil ke besar :")
	for i:= 0; i < n; i++ {
		fmt.Println(b[i].wilayah, b[i].jumLahir, b[i].beratRendah, b[i].giziKurang)
	}
}

func main() {
	var bayi tabBayi
	var n int

	fmt.Scan(&n)
	bacaData(&bayi, n)
	urutMembesar(&bayi, n)
	tampilArrayBayi(bayi, n)

}