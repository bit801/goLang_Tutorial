//	Go supports constants of character, string, boolean, and numeric values

package main

import "fmt"

//const declares a constant value
const s string = "constant"

func main() {
	fmt.Println(s)

	//A const statement can appear anywhere a var statement can
	const n = 500000000

}
