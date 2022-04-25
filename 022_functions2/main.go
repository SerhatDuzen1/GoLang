package main

import "fmt"

func main() {
	//1
	fmt.Println(split(17))

	//2
	numTerms, sum := add(1,3)
	println("Added:",numTerms,"terms for a total of", sum)

	//3 değişken sayıda Parametre alan fonksiyonlar
	sayHello("Merhaba","Go","Nasılsın")

	//4 
	result := add2(4,5,3,4,5,6,7,8)
	println(result)

	//anonim fonksiyonlar - func ismi bilinmeksizin ulaşmaya çalışılır
	//5
	addFunc := func(terms ...int) (numnumTerms2 int, sum2 int){
		for _,term := range terms{
			sum2+=term
		}
		numnumTerms2 = len(terms) 
		return
	}
	numTerms2, sum2 := addFunc(2,5)
	println("Added :",numTerms2,"terms fora total of",sum2)
}	


//1
func split(sum int) (x, y int) {
	x =sum*4/9
	y=sum-x
	return //burda x ve y yi otomatik olarak bulur ve döndürür
}

//2
func add(terms ...int) (numTerms int, sum int){
	for _,term := range terms{
		sum+=term
	}
	numTerms = len(terms)
	return
}

//3
func sayHello(messages ...string){ //burdaki 3 nokta değişken sayıda girdi almamızı sağlar
	for _, message := range messages{
		println( message)
	}

}

//4

func add2(terms ...int) int{
	result:=0
	for _,term:= range terms{
		result+=term
	}
	return result
}

//5
 
