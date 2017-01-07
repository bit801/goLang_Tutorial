//	Go by example: If/Else
//	This prrogram demonstrates branching with if/else

package main

import "fmt"

func main() {

	fmt.Println()

	if 7%2 == 0 { //if 7 can be divided by 2 with no remainder
		fmt.Println("7 is even.") //7 is even
	} else {
		fmt.Println("7 is odd.") //else, 7 divided by 2 has a remainder which
		//means 7 is odd
	}

	if 8%4 == 0 { //if 8 can be divided by 4 with no remainder
		fmt.Println("8 is divisible by 4.") //then 8 8 is divisible by 4
	}

	if num := 9; num < 0 { //syntax shorthand for declaring and initializing variable 'num' to 9
		//also infers the 'type' of initialized variable 'num' to int
		//if num is less than 0
		fmt.Println(num, " is negative.") //num is negative
	} else if num < 10 {
		fmt.Println(num, "has 1 digit.")
	} else {
		fmt.Println(num, " has multiple digits.")
	}

	fmt.Println()
}
