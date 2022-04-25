package main

import "fmt"

func main() {
	str1 := "Brad"
	str2 := "Ali"
	str3 := "Ayşe"
	aNumber := 20
	isTrue := true
	stringLenght, _ := fmt.Println(str1,str2,str3) //hem çıktı verir, = ,le de kaç karakter olduğunu int e atama yapılır

	//if err == nil {
		fmt.Println("String lenght:",stringLenght)
		//}
		fmt.Printf("Value of aNumber : %v\n",aNumber) // v placeholder dır 
		fmt.Printf("Value of isTrue : %v, mesela %v\n",isTrue,false) // birden fazla değer de gelebilir
		fmt.Printf("Value of aNumber as Float: %.2f\n",float64(aNumber)) // %.2f 

		//GOLANG Placeholder internetten araştırma yapılabilir

		fmt.Printf("Data types as var: %T, %T, %T, %T, and %T",str1,str2,str3,aNumber,isTrue)


		fmt.Println("***********************")

		handsome := "Brad"
		charizmatic := "O da Brad"
		excellent := "Tabiki Brad"
		age := 34.99
		fmt.Printf("Bu alemin en yakışıklısı %v\n",handsome)
		fmt.Printf("Bu alemin en karizması %v\n",charizmatic)
		fmt.Printf("Bu alemin en excellent ı %v\n",excellent)
		fmt.Printf("Onun yaşı da %2f\n",age)
}