package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

func main() {

	file, err := os.Open("dosyam.txt")
	// file, _ := os.Open("dosyam.txt") //err yi yönetmek istemiyorsan

	if err != nil {
		fmt.Println(err.Error())
	}

	// yukarıdakine alternatif olarak kısa çözüm 
	checkError(err)

	fmt.Println(file.Name)


	myError := errors.New("Bu bir hata")
	fmt.Println(myError.Error())

	// i := -2
	// if i<0 {
	// 	return 0, fmt.Errorf("matematik : çok kötü bir hata %g",i)
	// }

}

func checkError (err error) {
	if err != nil {
		fmt.Println("Error : ",err.Error())
		//logladınız
		log.Fatal(" ERROR : ",err)
		os.Exit(1)
	}
}