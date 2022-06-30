package main

import "fmt"

// global variable declaration
var global int = 10

func main() { // from here local level scope starts

	// local variables inside the main function
	var local int = 20

	// Display the value of global variable
	fmt.Printf("The value of Global variable is : %d\n",
		global)

	// Display the value of local variable
	fmt.Printf("The value of Local variable is : %d\n",
		local)

	display()

} // here local level scope ends

// taking a function
func display() { // local level starts

	// Display the value of global variable
	fmt.Printf("The value of Global variable is : %d\n",
		global)

} // local scope ends here
