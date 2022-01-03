package main

import (
	"fmt"
)

func doAWhileLoop(){
	var sum int = 0;
	for sum <= 10 {
		sum += 2;
		fmt.Println(sum)
	}
}

func doAForLoop(){
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func doLoopInArray(){
	arr := []string{"hello", "goodbye"};

	for i, str := range arr {
		fmt.Println(i);
		fmt.Println(fmt.Sprintf("%s Jimmy", str));
	}
}

func main(){
	doAForLoop();
	doAWhileLoop();
	doLoopInArray();
}
