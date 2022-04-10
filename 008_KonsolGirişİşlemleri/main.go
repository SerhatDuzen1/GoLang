package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	//konsol veri girişi
	reader := bufio.NewReader(os.Stdin)
	
	fmt.Print("Enter Text : ")
	str, _ := reader.ReadString('\n')
	fmt.Println(str)

	fmt.Println("Enter a number : ")
	str, _ = reader.ReadString('\n')
	f, err := strconv.ParseFloat(strings.TrimSpace(str), 64) // stringin sağında ve solundaki boşlukları trimler
	
	if err != nil {
		fmt.Println(err)
	} else{
		fmt.Println("Value of number : ", f)
	}


	reader2 := bufio.NewReader(os.Stdin)
	fmt.Println("yaz bişi")
	str2, _ :=reader2.ReadString('\n')
	fmt.Println(str2)

	reader3 := bufio.NewReader(os.Stdin)
	fmt.Println("sayı yaz")
	int2,_ := reader3.ReadString('\n')
	kk,_ := strconv.ParseFloat(strings.TrimSpace(int2),64)
	fmt.Println(kk)
}