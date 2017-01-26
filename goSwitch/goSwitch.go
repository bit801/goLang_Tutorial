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
	case 1: //must have colon after case: 1 go to case 1. Don't forget the colon after the case
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
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	//A "type switch" compares types instead of values. You can use this to discover the type
	//of an interface value. In this example, the variable 't' will have the type corresponding
	//to its clause.
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Println("Don't know the type %Tn", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")

}
