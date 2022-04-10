package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {

	//

	for i, v := range pow {  //i index v ise value
		fmt.Printf("2**%d = %d\n", i, v)
	}

// array örneği
	a := [...]string{"a","b","c","d","e"}

	for i := range a {
		fmt.Println("Array item ",i," is ",a[i])
	}

//map örneği

	baskentler1 := map[string]string{"Turkey":"Ankara", "France":"Paris", "Italy":"Rome", "Japan":"Tokyo", "Russia":"Moscow"}
	for key := range baskentler1 {
		fmt.Println("Map Item: Capital of", key,"is",baskentler1[key])
	}

	//diğer yöntem

	for k, val := range baskentler1 { // k index val ise value
		fmt.Println("Map Item: Capital of", k,"is",val)
	}
}