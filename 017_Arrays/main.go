package main

import (
	"fmt"
)

func main() {
	
	// var a [10]int
	// fmt.Println(a)

	// myArray1 := [3]int{}
	// myArray1[0] =32
	// myArray1[1] =23
	// myArray1[2] =54
	// fmt.Println(myArray1)

	// var colors [3]string
	// colors[0]="Red"
	// colors[1]="Green"
	// colors[2]="Blue"
	// fmt.Println(colors)
	// colors[1]="Black"
	// fmt.Println(colors[1])

	var numbers = [5]int {1,45,76,2,6}
	fmt.Println(numbers)
	fmt.Println("Number of numbers ",len(numbers)) // dizi uzunluğu
	
	// myArray2 := [4]int{4,3,6,7} //normali
	array2 := [...]int{3,4,5,6} //automatic size
	array2[3] = 555
	fmt.Println(array2)

	var cars [3] string
	cars[0]="Opel" 
	cars[1]="Reno" 
	cars[2]="Ferrari" 
	fmt.Println(len(cars))
	fmt.Println(cap(cars))

	i := 0
	for i<len(cars) {
		fmt.Println(cars[i])
		i++
	}

	for i := 0; i < len(cars); i++ {
		fmt.Println(cars[i],"\n")
	}

	for i,value := range cars {
		fmt.Println("i = ",i,"&value = ",value)
	}







}