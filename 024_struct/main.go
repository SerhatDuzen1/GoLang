package main

import (
	"fmt"
	"strconv"
)

//diğer dillerdeki class yerine struct kullanılır

type Human struct {
	FirstName string
	LastName  string
	Age       int
}

func newHuman() *Human { // varsayılan ve boş yapıcı metot /parametresiz
	h := new(Human)
	return h
}

func newHumanP(FirstName, LastName string, age int) *Human { //GoLang metot overload desteklemez
	r := new(Human)
	r.FirstName = FirstName
	r.LastName = LastName
	r.Age = age
	return r
}

func main() {
	//Nesne örneği oluşturma V1
	x := Human{FirstName: "Serhat"}
	fmt.Println(x.FirstName)

	//V2
	y := new(Human)
	y.Age = 25
	fmt.Println(y.Age)

	//Yapıcı metot

	z := newHuman()
	z.FirstName = "Fatih"
	fmt.Println(z.FirstName)

	b := newHumanP("Ilhan", "Mansiz", 45)
	fmt.Println(b.LastName)

	//Nesnenin verisini okumak
	var data = b.FirstName + " " + b.LastName + " " + strconv.Itoa(b.Age)
	fmt.Println(data)
}
