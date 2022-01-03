package main

import (
	"fmt"
	"importingPackages/hello"
)

func main(){
	fmt.Println(hello.ReturnAddMethod(300, 300))
	fmt.Println(hello.HelloWorld("Jimmy"))
}
