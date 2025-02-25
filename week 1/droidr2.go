//SOAL

// Droid R2-D2 mengirim pesan rahasia untuk C-3PO dalam bentuk bahasa mesin.
// Setelah diterjemahkan dan dicetak, maka akan tampak pola bintang seperti seperti dua segitiga siku-siku.

// Buatlah program untuk mencetak pola bintang itu.

// Masukan berupa bilangan asli positif ganjil.

// Keluaran pola bintang seperti tampak pada tabel

// Petunjuk: Gunakan perulangan bersarang (nested loop).

package main
import "fmt"

func main() {
    var x int
    fmt.Scan(&x)
    for i := 1; i <= x; i++{
        for j := 1; j <= x; j++{
            if(j != i){
                fmt.Print("*")
            }else{
                fmt.Print(" ")
            }
        }
        fmt.Println("")
    }
}