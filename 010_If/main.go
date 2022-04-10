package main

import "fmt"

func main() {
	//IF

//  a := 10
// 	b := 10
	// if b > a {
	// 	fmt.Println("Büyüktür")
	// }else if b==a {
	// 	fmt.Println("Eşittir")
	// }else {
	// 	fmt.Println("küçüktür")
	// }
// foo :=1
// if foo==1 {
// 	fmt.Println("bar")
// }
	if foo := 2; foo ==1 {
		fmt.Println("bar")
	}else {
		fmt.Println("buz")
	}
	// fmt.Println(foo) yazmaz çünkü local değişken

	yas := 35

	if yas ==35 {
		fmt.Println("Yolun yarısı eder")
	} else if (yas < 35) {
		fmt.Println("Yolun başındasın")
	} else {
		fmt.Println("Sona yaklaşmışsın be anam")
	}
}