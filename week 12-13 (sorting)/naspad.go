package main
import "fmt"

const NMAX int = 1024

type naspad struct {
	nama string
	rating, harga int
}
type tabNaspad [NMAX]naspad

func bacaNaspad(np *tabNaspad, n int) {
	for i := 0; i < n; i++ {
		fmt.Scan(&np[i].nama, &np[i].rating, &np[i].harga)
	}
}

func urut(np *tabNaspad, n int){
	var pass, idx int
	var temp naspad
	for pass = 0; pass < n-1; pass++ {
		idx = pass
		for j := pass + 1; j < n; j++ {
			if np[idx].rating < np[j].rating {
				idx = j
			} else if np[idx].rating == np[j].rating && np[idx].harga > np[j].harga {
				idx = j
			}
		}
		temp = np[idx]
		np[idx] = np[pass]
		np[pass] = temp
	}

}

func cetakNaspad(np tabNaspad, n int) {
	fmt.Println("mengurutkan rating besar dan harga murah")
	for i:=0; i<n; i++ {
		fmt.Println(np[i].nama, np[i].rating, np[i].harga)
	}
}

func main () {
	var padang tabNaspad
	var n int

	fmt.Scan(&n)
	bacaNaspad(&padang, n)
	urut(&padang, n)
	cetakNaspad(padang, n)

}