package main

import "fmt"

func main() {

	//The most basic type of for loop with a single condition
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + i
	}
}
