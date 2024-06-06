const nMax int = 1000
type tabInt [nMax]int
func isiArray(arr1, arr2 *tabInt, n int) {
	/* I.S. Data tersedia dalam piranti masukan
	F.S. arr1, arr2 telah terisi sejumlah bilangan */
	var ... // variabel digunakan dalam perulangan dan  menyimpan sementara nilai untuk masukkan
    ...
	for i:= 0; i < n; i++{ // mengisi array 1
		fmt.Scan(&input)
		for input != 0 && input != 1 { // menerima inputan hingga inputan yang dimasukkan 1 atau 0
			fmt.Scan(&input)
		}
		arr1[i] = input
	}

	for i < n { // mengisi array 2
		fmt.Scan(&input)
		for input != 0 && input != 1 { // menerima inputan hingga inputan yang dimasukkan 1 atau 0
			fmt.Scan(&input)
		}
		arr2[i] = input	
	}
}

func and(x1, x2 int) int {
	/*mengembalikan nilai 1 apabila memenuhi logika and dan 0 jika tidak memenuhi*/
	if x1 == 1 && x2 == 1{
		return 1
	} else {
		return 0
	}
}

func or(x1, x2 int) int {
	/*mengembalikan nilai 1 apabila memenuhi logika OR dan 0 jika tidak memenuhi*/
	if x1 == 1 || x2 == 1{
		return 1
	} else {
		return 0
	}
}

func hitungOR(arr1, arr2 tabInt, n int) int {
	/*mengembalikan total dari banyaknya operasi or yang memenuhi kondisi pada array*/
	var total int // variabel digunakan untuk iterasi loop dan menyimpan total banyaknya operasi yang memenuhi
	for i := 0; i < n; i++ { // iterasi sebanyak n kali
		val := or(arr1[i], arr2[i]) // menghitung operasi or pada elemen array
		if val == 1 { // jika memenuhi kondisi maka total dari jumlah operasi yang memenuhi bertambah
			total++
		}
	}
	return total
}

func hitungAND(arr1, arr2 tabInt, n int) int {
	/*mengembalikan total dari banyaknya operasi logika and yang memenuhi kondisi pada array*/
    var total int
	for i := 0; i < n; i++ {
		val := and(arr1[i], arr2[i])
		if val == 1 {
			total++
		}
	}
	return total
}

func xor(x1, x2 int) int {
	/*mengembalikan nilai 1 apabila memenuhi logika XOR dan 0 jika tidak memenuhi*/
	if x1 != x2 { // syarat ketika memenuhi XOR yaitu ketika nilai x1 != nilai x2
		return 1
	} else {
		return 0
	}
}

func hitungXOR(arr1, arr2 tabInt, n int) int {
	/*mengembalikan total dari banyaknya operasi xor yang memenuhi kondisi pada array*/
	var total int
	for i := 0; i < n; i++ {
		val := xor(arr1[i], arr2[i])
		if val == 1 {
			total++
		}
	}
	return total
}

func cetak(arr1, arr2 tabInt, n int) {
	/* I.S. array arr1,arr2 berisi sejumlah n bilangan 1 dan 0
	F.S. Banyaknya hasil dari operasi logika OR, hasil dari operasi logika OR, banyaknya hasil dari operasi logika AND
	hasil dari operasi logika AND, banyaknya hasil dari operasi logika XOR,hasil dari operasi logika OR ditampilkan di layar */
	var i, totalOR,...,...,... int

	fmt.Println("Total OR:", totalOR) // menampilkan total OR yang memenuhi
	for i = 0; i < n; i++ {
		fmt.Print(or(arr1[i], arr2[i]," "))// menampilkan hasil hasil dari operasi logika OR(perhatikan contoh !!!)
	}
	fmt.Println()

	fmt.Println("Total XOR:", ....) // menampilkan total operasi logika XOR yang memenuhi
	for i y= 0; i < n; i++{
		...
		...// menampilkan hasil hasil dari operasi logika XOR(perhatikan contoh !!!)
	}
	fmt.Println()
}
