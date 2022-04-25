package main



func main() {

// 	message := "Hi"
// 	sayHello(&message)
// 	fmt.Println(message)
// }

// func sayHello(message *string) {
// 	println(*message)
// 	*message = "hi GO!"
// }

// fmt.Println(add(1065,567))

// }

// func add(x , y int) int { //en sağdaki int soldakileri niteler
// 	return x+y
// }

// a,b := swap("ali ata","bak")
// fmt.Println(a,b)
// }

// //çoklu sonuç dönen func,

// func swap(x,y string) (string, string) {
// 	return y,x 
// }

numTerms , sum := add(1,2,3,4,5,6)
println("Added :",numTerms,"terms for a total of",sum)
}

func add(terms ...int) (int, int) { //çoklu sonuç ve girdi
	result := 0
	for _,term := range terms {
		result += term
	}
	return len(terms), result
}