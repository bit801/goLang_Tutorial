//	Go by Example: Switch
//	Switch statements express conditionals across many branches
package main

import "fmt"

func main() {

	//this is a basic switch
	i := 1
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}
}
