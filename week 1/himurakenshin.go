//SOAL

// Untuk mencegah penghancuran kota Kyoto, Himura Kenshin harus 
// menghadapi Shishio Makoto yang memiliki anak buah 
// sebanyak n samurai dengan beragam kekuatan dan kecepatan. 
// Menyadari kalah dalam jumlah, Kenshin menerapkan strategi berikut agar dapat mengalahkan sebanyak-banyaknya lawan: 

// *Kenshin akan melakukan pertarungannya untuk seluruh lawan dan selama kekuatan dan kecepatannya minimal 3.

// *Setiap mengalahkan lawan dengan menggunakan kekuatan atau kecepatan, maka kekuatan atau kecepatan Kenshin berkurang 6.

// *Jika lawan yang dihadapi memiliki kekuatan dan kecepatan yang lebih besar dari atau sama dengan Kenshin, maka Kenshin akan menghindarinya.

// *Jika lawan yang dihadapi memiliki kekuatan yang lebih besar dari atau sama dengan Kenshin, 
// namun kecepatan yang lebih kecil dari Kenshin, maka Kenshin akan menggunakan kecepatan untuk mengalahkannya.

// *Jika lawan yang dihadapi memiliki kekuatan yang lebih kecil, namun kecepatan yang lebih besar dari Kenshin, 
// maka Kenshin akan menggunakan kekuatan untuk mengalahkannya.

// *Jika lawan yang dihadapi memiliki kekuatan dan kecepatan yang lebih kecil dari Kenshin, 
// maka Kenshin akan menggunakan salah satu dari sisa yang lebih besar antara kekuatan atau kecepatannya; 
// atau kecepatannya jika sisa kekuatan sama dengan kecepatan.

// Buatlah program yang mengimplementasikan pertarungan Kenshin itu.

// Masukan berupa banyaknya n samurai yang harus dihadapi Kenshin. kekuatan dan kecepatan Kenshin saat ini dalam bilangan bulat. Sebanyak n baris pasangan bilangan bulat yang menyatakan kekuatan dan kecepatan lawan.

// Keluaran berupa jumlah lawan yang berhasil dikalahkan dan sisa kekuatan serta kecepatan Kenshin.

package main

import "fmt"

func main() {
	var n_kali, kec_kenshin, ket_kenshin, kec_musuh, ket_musuh, menang int
	fmt.Scan(&n_kali)
	fmt.Scan(&ket_kenshin, &kec_kenshin)
	for i := 0; i < n_kali; i++{
		fmt.Scan(&ket_musuh, &kec_musuh)
		if kec_kenshin >=3 && ket_kenshin >=3 {
			if kec_kenshin <= kec_musuh && ket_kenshin <= ket_musuh{
				continue
			} else if ket_kenshin <= ket_musuh && kec_kenshin > kec_musuh{
				menang++
				kec_kenshin-=6
			} else if ket_kenshin > ket_musuh && kec_kenshin <= kec_musuh{
				menang++
				ket_kenshin-=6
			} else if ket_kenshin > ket_musuh && kec_kenshin > kec_musuh{
				menang++
				kec_kenshin-=6
			} else if ket_kenshin <= ket_musuh && kec_kenshin > kec_musuh{
				menang++
				if ket_kenshin > kec_kenshin{
					ket_kenshin-=6
				} else{
					kec_kenshin-=6
				}
			}
		} else{
			continue
		}
	}
	fmt.Print(menang, ket_kenshin, kec_kenshin)
}