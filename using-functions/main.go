package main

import (
	"fmt"
)

func greetTheUser(user string) string {
	userPrint := fmt.Sprintf("Hello %s", user);
	return userPrint;
}

func doSum(num1 int, num2 int) int {
	return num1 + num2;
}

func doFloatMultiplication(num1 float32, num2 float32) float32 {
	return num1 * num2;
}

func main(){
	fmt.Println(greetTheUser("Jose"))
	fmt.Println(doSum(100, 100))
	fmt.Println(doFloatMultiplication(25.5, 9.2))
}
