package main

import "fmt"

func main() {

	//the most basic type of for loop with a single condition
	fmt.Println() //print blank line to screen
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + i
	}
	fmt.Println() //print blank line to screen

	//a classic initial/condition/after for loop
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}
	fmt.Println() //print blank line to screen

	//for loop without a conditon will loop repeatedly until you break out of the loop
	//or return from the enclosed function
	for {
		fmt.Println("loop without condition.")
		break
	}
	fmt.Println() //print blank line to screen

	//you can also continue to the next iteration of the loop
	for n := 0; n <= 5; n++ {
		if n%2 == 0 { //if index is even and less or equal to 5
			fmt.Println("Inside of for-loop at iteration: ") //continue to next iteration inside loop
			fmt.Println(n)
			fmt.Println()
			continue
		}
		fmt.Println("Outside of for-loop at iteration: ")
		fmt.Println(n)
		fmt.Println()
	}
}
