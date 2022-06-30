package main

import "fmt"

func main() {

	var a int = 100
	var b int = 45
	var c int

	c = a + b
	fmt.Printf("Line 1 - Value of c is %d\n", c)

	c = a - b
	fmt.Printf("Line 2 - Value of c is %d\n", c)

	c = a * b
	fmt.Printf("Line 3 - Value of c is %d\n", c)

	c = a / b
	fmt.Printf("Line 4 - Value of c is %d\n", c)

	c = a % b
	fmt.Printf("Line 5 - Value of c is %d\n", c)

}
