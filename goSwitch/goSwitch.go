//	Go by Example: Switch
//	Switch statements express conditionals across many branches
package main

import "fmt"
import "time"

func main() {

	//This is a basic switch
	//the := sytax is shorthand for declaring and initializing a variable,
	//e.g. for var i int = 1 in this case
	i := 1
	fmt.Print("Write ", i, " as ")
	switch i { //switch, if i is equal to
	case 1: //must have colon after case: 1 go to case 1
		fmt.Println("one")
	case 2: //2 go to case 2
		fmt.Println("two")
	case 3: //3 go to case 3
		fmt.Println("three")
	}

	//You can use commas to separate multiple expressions in
	//the same case statement. We use the optional 'default'
	//case in this example as well
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday: //comma used to separate expressions
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	//Switch without an expression is an alternate way to express if/else logic.
	//Here we also show how the case expressions can be non-constants

}
