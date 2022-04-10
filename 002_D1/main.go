package main

import "fmt"

const aciklama = "Go Uygulaması"//başlangıçta verilen sabit değerler
const pi =3.14 

var a = "ABC DSA" //global alan
var d string = "Serhat"

var (
	degisken1 = "Serhat"
	degisken2 = "BJK"
)


func main() {
	/*
	var message string
	message = "Merhaba Go!"
	fmt.Println(message)

	var message = "Merhaba!"
fmt.Println(message)

var a, b, c int = 3, 4, 5 
fmt.Println(a)  //default değerleri 0 dır
fmt.Println(b)
fmt.Println(c)


var message = "Hello Go!"
var a, b, c = 3, true, 4.5 // çoklu atama da yapılabilir
fmt.Println(message,a,b,c)

u := 55 // kısa yol
v, n := "abc", true
message := "Merhaba Go!"

a := "Metin"
b := 'M' //ascii karşılığını verir 77,
fmt.Println(a, b) 

sayi := 4
fmt.Println(sayi)
aciklama := "asdsdafsa"
fmt.Println(aciklama)

*/




}

// sayi := 4 // hata verir, bir fuction un dışında yani global alanda var keywordu kullanmak gereklidir
var name string = "golang"   //private variable
var Version string = "1.2.3" //public variable
