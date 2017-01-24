//	Go by Example: Switch
//	Switch statements express conditionals across many branches
package main

import "fmt"

func main() {

	//This is a basic switch
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

//You can use commas to separate multiple expressions in
//the same case statement. We use the optional 'default'
//case in this example as well
