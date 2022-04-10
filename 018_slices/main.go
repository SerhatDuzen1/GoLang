package main

import (
	"fmt"
)

func main() {

	primes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 31, 37}

	var s []int = primes[1:4] // Dilimler diziliere olan başvurular gibidir

	fmt.Println(primes)

	fmt.Println(s)

	primes[2] = 4

	fmt.Println(s)

	s[1] = 5

	fmt.Println(primes)

	fmt.Println(primes[2:]) //ikinci indexden dilim uzunluğuna kadar git

	fmt.Println(primes[:5]) //sıfırıncı indexden beşinci indexe kadar git

	fmt.Println("------------------------------------")

	myArray := [...]int{45, 23, 43}
	mySlice := myArray[:]
	fmt.Println(mySlice)
	mySlice[0] = 11
	fmt.Println(mySlice)
	fmt.Println(myArray)

	fmt.Println("------------------------------------")

	var colors = []string{"red", "green", "blue"}
	fmt.Println(colors)
	colors = append(colors, "purple")
	fmt.Println(colors)

	colors = append(colors[1:len(colors)])
	fmt.Println(colors)
	colors = append(colors[:len(colors)-1])
	fmt.Println(colors)

	fmt.Println("------------------------------------")

	numbers := make([]int, 5, 5)
	numbers[0] = 123
	numbers[1] = 43
	numbers[2] = 36
	numbers[3] = 4
	numbers[4] = 435
	fmt.Println(numbers)
	numbers = append(numbers, 765)
	fmt.Println(numbers)
	fmt.Println(len(numbers))
	fmt.Println(cap(numbers))

	fmt.Println("------------------------------------")

	//var newSlice = make ([]int,5,5)
	var slice []int
	for i := 0; i < 10; i++ {
		k := i + 1
		var j int
		fmt.Printf("Lütfen %v'inci sayıyı giriniz : ",k )
		fmt.Scan(&j)
		slice = append(slice, j)
		
	}
	fmt.Println(slice)
	var y string
	fmt.Println("lütfen isminizi giriniz")
	fmt.Scan(&y)
	fmt.Printf("Hoşgeldiniz Sn. %v",y)
}
