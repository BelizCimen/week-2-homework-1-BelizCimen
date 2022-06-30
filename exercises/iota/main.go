package main

import "fmt"

const (
	Sunday int = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func main() {
	fmt.Printf("The value of Red    is %v\n", Sunday)
	fmt.Printf("The value of Orange is %v\n", Monday)
	fmt.Printf("The value of Yellow is %v\n", Tuesday)
	fmt.Printf("The value of Green  is %v\n", Wednesday)
	fmt.Printf("The value of Blue   is %v\n", Thursday)
	fmt.Printf("The value of Indigo is %v\n", Friday)
	fmt.Printf("The value of Violet is %v\n", Saturday)
}
