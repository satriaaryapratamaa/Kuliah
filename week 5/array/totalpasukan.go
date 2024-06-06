package main
import "fmt"

type pasukan struct {
	nama string
	artileri, kavaleri int
}

func total(x,y,z pasukan) int{
	return x.artileri + x.kavaleri + y.artileri + y.kavaleri + z.artileri + z.kavaleri
}

func main() {
	var a, b, c pasukan
	
	fmt.Scan(&a.nama, &a.artileri, &a.kavaleri)
	fmt.Scan(&b.nama, &b.artileri, &b.kavaleri)
	fmt.Scan(&c.nama, &c.artileri, &c.kavaleri)
	
	fmt.Println(total(a,b,c))
}
