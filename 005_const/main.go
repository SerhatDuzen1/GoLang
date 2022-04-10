package main

import "fmt"

const Pi = 3.14

const thy = "Turkish Airlines"

const tai = "TAI"

const (
	BJK Brand =  "Beşiktaş"
	MS  Brand =  "Microsoft"
)

func PrintBrand(brand Brand){
fmt.Println(brand)
}

type Brand string  // bu şekilde bir tip oluşturulabilir

func main() {
	
	PrintBrand(BJK)
	PrintBrand(MS)
	// thy = "abc" const onceden tanımlandığı için hata fırlatır
	fmt.Println(thy,tai)



}