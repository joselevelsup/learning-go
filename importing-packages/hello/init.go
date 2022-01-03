package hello

import (
	"fmt"
)

func doMath(num1 int, num2 int) int {
	return num1 + num2;
}

func ReturnAddMethod(num1 int, num2 int) int {
	var sum int = doMath(num1, num2)
	return sum
}

//private functions start with a lowercase letter while a public function would be in a capitalized
func HelloWorld(name string) string {
	return fmt.Sprintf("Hello %s", name);
}
