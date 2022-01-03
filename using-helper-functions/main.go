package main

import (
	"fmt"
	"github.com/thoas/go-funk" //if lsp tells you there is an error with this import like "no required module provides package" its fine. go run main.go works just fine
)

func printMap(){
	// maps are just like how objects are in javascript
	sampleMap := map[string]int{"Mark": 10, "Jimmy": 20}
	fmt.Println(sampleMap);
}

func printArray(){
	//Can also declare a max of items in array
	// x := [2]int === [0, 0]
	arr := []string{"hello", "goodbye"}

	fmt.Println(arr)
}

func mapThroughArray(){
	arr := [4]int{};

	result := funk.Map(arr, func(x int) int {
		return x + 1
	});

	fmt.Println(result);
}

func main(){
	mapThroughArray();
}
