package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	num := 10
	str := "20"

	//number to string
	fmt.Println(num, reflect.TypeOf(num))
	numStr := strconv.Itoa(num)
	fmt.Println(numStr, reflect.TypeOf(numStr))

	//string to number
	fmt.Println(str, reflect.TypeOf(str))
	strNum, err := strconv.Atoi(str)
	if err != nil {
		return
	}
	fmt.Println(strNum, reflect.TypeOf(strNum))

}
