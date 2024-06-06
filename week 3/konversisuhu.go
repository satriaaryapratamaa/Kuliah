// Buatlah program(dalam bentuk pseudocode) yang mengimplementasikan sebuah 
// prosedur untuk mengkonversi temperatur dalam satuan Celsius (C) ke dalam satuan 
// Reaumur (R), Fahrenheit (F) dan Kelvin (K)! 
// Berikut adalah perhitungan konversinya C=5/9(F−32), C=5/4 R, dan K=C+273.15 
// Masukan terdiri dari bilangan riil yang menyatakan temperatur dalam satuan Celsius. 
// Keluaran adalah Ɵga bilangan riil yang menyatakan hasil konversi Celsius dalam satuan 
// Reaumur, Fahrenheit dan Kelvin. 

package main

import "fmt"

func pilih_celcius(x string) {
	fmt.Println("Ini adalah Program konversi celcius")
	fmt.Println("Masukkan Angka dalam celcius : ")
}

func program_konversi_celcius(z float64) {
	var reamur, fahrenheit, kelvin float64

	reamur = (z *4) / 5 
	fahrenheit = (z * 9 / 5) + 32
	kelvin = z + 273.15

	fmt.Print(reamur,"R ")
	fmt.Print(fahrenheit,"F  ")
	fmt.Print(kelvin,"K")

}

func main() {
	var C float64
	var x string
	fmt.Println("Sudah Siapkah kamu?")

	fmt.Scan(&x)
	pilih_celcius(x)
	fmt.Scan(&C)
	program_konversi_celcius(C)

}
