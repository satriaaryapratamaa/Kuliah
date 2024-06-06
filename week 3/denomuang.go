package main
import "fmt"

func main() {
	var u int
	var u10,u2,u1 int

	fmt.Scan(&u)
	pecahUang(u,&u10,&u2,&u1)
	fmt.Println(u10,u2,u1)
}

func pecahUang(uang int, k10,k2,k1 *int) {
	*k10 = uang/10000
	uang = uang%10000

	*k2 = uang/2000
	uang = uang%2000

	*k1 = uang/1000
}
