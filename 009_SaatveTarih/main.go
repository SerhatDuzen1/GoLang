package main

import (
	"fmt"
	"time"
)

//


func main() {
	// saat ve tarih işlemleri
	//date() metodu ile kendi tarih verimizi oluşturuyoruz
	t := time.Date(2016, time.April, 10, 20, 0 ,0, 0, time.UTC ) 

	//T ADIYLA tarih tipinde oluşturulan veriyi string sipinde ekrana yazıyoruz
	fmt.Printf("Go launch at %s\n", t)

	fmt.Println("**************")

	now := time.Now()

	fmt.Printf("The time now is %s\n ",now)
	fmt.Println("************")

	fmt.Println(t.Month())
	fmt.Println(t.Day())
	fmt.Println(t.Weekday())

	tomorrow := t.AddDate(0,0,1)
	fmt.Printf("tomorrow is %v %v %v %v\n",tomorrow.Weekday(),tomorrow.Month(),tomorrow.Day(),tomorrow.Year())


	t2 := time.Date(1987, time.April, 6,6,0,0,0,time.Local)
	fmt.Printf("Serhat ın doğum tarihi %s\n",t2)
	now2 := time.Now()
	fmt.Println(now2)
	nextmonth := now2.AddDate(1,2,3)
	fmt.Println(nextmonth)
}

// END OMIT
