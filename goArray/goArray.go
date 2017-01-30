//In Go, an array is a numbered sequence of elements of a specific length

package main

import "fmt"

func main() {
	//Here we create an array that will hold exactly 5 ints. The type of elements and length are both part of
	//the array's type. By default an array is zero-valued, which for ints means 0's
	var a [5]int //declare array named 'a' of type int with 5 elements

	a[4] = 100 //set value at index 4 to 100
	//basically set value at index using array[index]=value syntax
	fmt.Println("set:", a)
	fmt.Println("get:", a[4]) //we get a value with array=[index] syntax

	//the built in 'len' returns the length of an array
	fmt.Println("len:", len(a))

	//use this syntax to declare and initialize an array in one line
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	//array types are one-dimensional, but you can compose types to build
	//multi-dimensional data structures
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		fmt.Println("outer for-loop", i)
		for j := 0; j < 3; j++ {
			fmt.Println("inner for-loop", j)
			twoD[i][j] = i + j
			fmt.Println("2d: ", twoD)
		}
	}
	fmt.Println("2d: ", twoD)

}
