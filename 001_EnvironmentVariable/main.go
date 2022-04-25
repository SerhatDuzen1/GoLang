package main

import (
	"fmt"
	"os"
)

func main() {
	//ENVIRONMENT VARIABLE LISTELEME
	// for _, env := range os.Environ() {
	// 	fmt.Println(env)
	// }

//export GOPATH=$HOME/go
//export PATH=$PATH:$GOROOT/bin:$GOPATH/bin

	uName := os.Getenv("USERNAME")
	uDomain := os.Getenv("USERDOMAIN")
	processorArchitecture := os.Getenv("PROCESSOR_ARCHITECTURE")
	//processorIdentifier := os.Getenv("PROCESSOR_IDENTIFIER")
	processorLevel := os.Getenv("GOROOT")
	goPath := os.Getenv("GOPATH")
	goRoot := os.Getenv("GOROOT")
	//homePath := os.Getenv("HOMEPATH")


	fmt.Println("Username : "+ uName)
	fmt.Println("Domain : "+ uDomain)
	fmt.Println("İşlemci Mimarisi : "+ processorArchitecture)
	fmt.Println("İşlemci seviyesi : "+ processorLevel)
	fmt.Println("GOPATH : "+ goPath)
	fmt.Println("GOROOT : "+ goRoot)

}
