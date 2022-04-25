package main

import(
	"fmt"
	
)
func main() {
	//map oluşturma 1. yntem

	myMap := make(map[int]string)
	fmt.Println(myMap)
	myMap[43]="first"
	myMap[66]="second"
	fmt.Println(myMap) //key-value şeklinde basar

	//map oluşturma 2. yöntem
	states := make(map[string]string)
	states["CAL"]="California"
	states["PAR"]="Paris"
	states["LON"]="London"
	states["IST"]="Istanbul"
	fmt.Println(states)

	fmt.Println(states["PAR"]) //par key ine sahip value yi yazdırma

	delete(states,"PAR") //par key ine sahip value silme

	fmt.Println(states)

	states["ROM"]="Rome" // map e değer ekleme

	for k,v := range states{ //yazdırma
		fmt.Printf("%v : %v\n",k,v)
	}

	//states map nesnesinin elemanı adedince kapasiteli bir key listesi oluştur

	keys := make([]string,len(states))

	i :=0
	for k := range states{
		keys[i]=k
		i++
	}
	fmt.Println(keys)

	//key listesindeki index değerlerine göre states nesnesinde bulunan şehirleri yazdır

	for i := range keys{
		fmt.Println(states[keys[i]])
	}
}