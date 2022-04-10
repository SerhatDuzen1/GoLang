package main

import (
	"fmt"
	"strconv"
)

func main() {
	// CONVERT
	var myString = "1"
	var x = 10
	var f float32 = 2.0
	fmt.Println(myString,x,f)
	//stringden int'e dönüştürme
	number, _ := strconv.Atoi(myString) //bu metot string to iint, mult,ple result durumu için _ konudu, error buraya gidecek

	fmt.Println(number+1) // artık int a dönüştüğü için toplama işlemine girer

	result := number*7

	//fmt.Println("Sonuç : "+result) HATA PATLATTI
	fmt.Println("Sonuç : "+ strconv.Itoa(result)) //Itoa string to int yapar

	// CASTING
	var i int = 55
	//var f1 float64 = i // büyükten küçüğe ya da küçükten büyüğe castingde metodunu kullanacaksın
	var fy float64 = float64(i)
	var u uint = uint(fy)
	fmt.Println(i,fy,u)


	str2 := "2354"
	int2 := 33
	flo2 := 3.3
	fmt.Println(str2,int2,flo2)
	number2 ,_ :=strconv.Atoi(str2)
	fmt.Println(number2) 
}