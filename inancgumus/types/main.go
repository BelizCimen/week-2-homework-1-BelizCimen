package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	type (
		Feet   float64
		Meters float64
	)

	var (
		feet   Feet
		meters Meters
	)

	arg := os.Args[1]

	// feet is a Feet value now
	val, _ := strconv.ParseFloat(arg, 64)
	feet = Feet(val)

	// meters is a Meters value now
	meters = Meters(feet * 0.3048)

	fmt.Printf("%g feet is %g meters.\n", feet, meters)
}
