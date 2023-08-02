package main

// https://pkg.go.dev/fmt

import (
	"fmt"
)

type X struct {
	Field1 string
	Field2 string
}

var (
	s                    = "hello!"
	ui16         uint16  = 1
	largeFloat   float64 = 12342234.234
	float        float64 = 3.14
	trueOrFalse          = 1 > 2
	x255                 = 255
	ASCII                = 65
	Percentage           = 42.1
	pointer      *X
	pointerValue = new(X)
	pointerSlice = make([]*X, 0)
	pointerMap   = make(map[string]*X)
)

func init() {
	pointerSlice = append(pointerSlice, &X{Field1: "value!", Field2: "value2!"})
}

func main() {

	fmt.Printf("Named Parameters: %(name)s is %(age)d years old\n", "name", "John", "age", 30)
	fmt.Printf("String: %s!\n", s)
	fmt.Printf("String Left Justify: %-6s!\n", s)
	fmt.Printf("String Right Justify: %6s!\n", s)

	fmt.Printf("Decimal: %d \n", ui16)
	fmt.Printf("True Or False?: %t \n", trueOrFalse)
	fmt.Printf("Left shift: %-6d \n", ui16)
	fmt.Printf("Right shift: %6d \n", ui16)
	fmt.Printf("Right shift with zero fill: %06d \n", ui16)
	fmt.Printf("Float with precision: %0.2f \n", float)
	fmt.Printf("Scientific Notation: %e\n", largeFloat)
	fmt.Printf("Hex: %x, Octal: %o\n", x255, x255)
	fmt.Printf("ASCII value: %c\n", ASCII)
	fmt.Printf("Percentage: %.2f%%\n", Percentage)

	fmt.Printf("Type: %T \n", ui16)
	fmt.Printf("Pointer Address: %p \n", &pointer)
	fmt.Printf("Pointer Value: %#v \n", *pointerValue)
	fmt.Printf("Map Values: %#v \n", pointerMap)
	fmt.Printf("Pointer Slice Values: %+q \n", pointerSlice)
	fmt.Printf("Pointer Slice index: %[1]q \n", pointerSlice)

	// MORE EXAMPLES CAN BE FOUND AT: https://pkg.go.dev/fmt
}
