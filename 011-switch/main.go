package main

import "fmt"

func main() {
	//switch

	foo := 3
	switch {
	case foo == 1:
		println("one")
	case foo == 2:
		println("two")
	case foo > 2:
		println("three")
	default:
		println("other")
	}

	var score float64
	fmt.Println("Enter score for your last exam : ")
	fmt.Scanf("%v", &score)
	switch{
	case score <=59:
		fmt.Println("Your grade is F")
	case score <=69:
		fmt.Println("Your grade is D")
	case score <=79:
		fmt.Println("Your grade is C")
	case score <=89:
		fmt.Println("Your grade is B")
	case score <=100:
		fmt.Println("Your grade is A")
	default:
		fmt.Println("Please enter a score <= 100")
	}


	var not2 float64
	fmt.Println("Yaz bekem son sınavdan ne alığını")
	fmt.Scanf("%v", &not2)
	switch{
	case not2 < 50:
		fmt.Println("...i tuttun")
	case not2 == 50:
		fmt.Println("ucu ucuna geçtin")
	case not2 > 50:
		fmt.Println("elal olsun rahat rahat geçtin")
	}
}