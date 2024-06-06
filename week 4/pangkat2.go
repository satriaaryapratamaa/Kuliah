// Buatlah program yang mengimplementasikan rekursif untuk mencari hasil pangkat dari dua buah 
// bilangan. 
// Masukan terdiri dari bilangan bulat x dan y. 
// Keluaran terdiri dari hasil x dipangkatkan y. 

package main
import "fmt"

func pangkat(x, y int) int {
	if x == 0 {
		return 0
	} else if y == 0{
		return 1
	}else {
		return x * pangkat(x, y-1)
	}
}

func main() {
	var x, y int

	fmt.Scan(&x, &y)
	fmt.Print(pangkat(x, y))
}