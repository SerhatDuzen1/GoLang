package main

import "fmt"

func main() {

	fmt.Println(" Şampiyon")
	a := 10
	b := 20
	
	total := a+b
	fmt.Println(total)

	total *=5
	fmt.Println((total))

	total %= 7
	fmt.Println("total : ",total)
}