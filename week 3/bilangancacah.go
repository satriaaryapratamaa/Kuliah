// Buatlah prosedur dalam bahasa Go bernama cacah untuk mencacah digit dari suatu bilangan bulat. 

// Masukan terdiri dari sebuah bilangan bulat positif X.

// Keluaran adalah beberapa nilai yang menyatakan nilai setiap digit dari X. Setiap nilai dipisahkan oleh spasi. Perhatikan contoh yang diberikan.

package main 
import "fmt"

func cacah(bilangan int) {
	/* I.S. terdefinisi sebuah bilangan bulat "bilangan"
	F.S. program mengeluarkan nilai yang menyatakan nilai setiap digit dari X dengan setiap nilai dipisahkan oleh spasi */
	for bilangan != 0 {
		fmt.Printf("%d ", bilangan%10)
		bilangan /= 10
	}
}

func main () {
	var x int

	fmt.Scan(&x)
	cacah(x)
}