package main
import "fmt"

type book struct {
	title string
	author string
	pages int
	price float64
	genre[4] string
}

func main (){
	book1 := book {
		title : "asimilikit goes haji",
		author : "Wahyu", 
		pages : 144,
		price : 1,
		genre : [4]string{"horor", "romance", "comedy", "crazy"},
	}

	// book2 := book {
	// 	title : "rahasia mokel tidak ketahuan",
	// 	author : "Udin", 
	// 	pages : 134,
	// 	price : 2000,
	// 	genre : [4]string{"horor", "romance", "comedy", "crazy"},
	// }

	// book3 := book {
	// 	title : "mencari wangsit",
	// 	author : "Udin", 
	// 	pages : 1,
	// 	price : 100000,
	// 	genre : [4]string{"horor", "romance", "comedy", "crazy"},
	// }

	fmt.Println("Buku 1")
	fmt.Println("Tittle", book1.title)
	fmt.Println("Author", book1.author)
	fmt.Println("Pages", book1.pages)
	fmt.Println("Price", book1.price)
	//fmt.Println("Kategori", book1.genre)

	for i := 0; i < len(book1.genre); i++ {
		fmt.Println("Kategori :", book1.genre[i])
	}


}