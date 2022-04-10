package main

import "fmt"

func main() {
	//Boş tanimlayici (_)

	// _, veri := metot() //Eğer burda hata şeklinde bir var tanimlamiş ve kullanmamiş olsaydik hata verecekti. blank var tanimlanabilir
	//fmt.Println(veri)

	//var _ string = "Serhat"
	//fmt.Println(_) böyle direk kullanilmaz

	var _, x, _ int=  0,9,0
	fmt.Println(x)
	strlenght, _ := fmt.Println("Nasilsin iyi misin burasinin uzunluğu kaç karakter.")
	fmt.Println(strlenght)
}