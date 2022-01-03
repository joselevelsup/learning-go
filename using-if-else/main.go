package main

import (
	"fmt"
)

// using Switch/case
// If a case is multiple lines of code, it needs to be in a bracket
func showDayOfTheWeek() {
	var dayOfWeek int = 6;
	switch dayOfWeek {
		case 1: fmt.Println("Monday")
		case 2: fmt.Println("Tuesday")
		case 3: fmt.Println("Wednesday")
		case 4: fmt.Println("Thursday")
		case 5: fmt.Println("Friday")
		case 6: {
			fmt.Println("Saturday")
			fmt.Println("Weekend. Yaay!")
		}
		case 7: {
			fmt.Println("Sunday")
			fmt.Println("Weekend. Yaay!")
		}
		default: fmt.Println("Invalid day")
	}
}

// using if else if statement and readline 
func main(){
	var color string;
	showDayOfTheWeek();

	fmt.Println("Enter a color");
	fmt.Scanf("%s", &color);
	
	if color == "red" {
		fmt.Println("you chose a redish color")
	} else if color == "blue" {
		fmt.Println("you chose a blueish color")
	} else {
		fmt.Println(fmt.Sprintf("Idk what this %s color is but get out", color));
	}
}
