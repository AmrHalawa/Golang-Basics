package main

import "fmt"

func main() {

	fmt.Println("Hello World") // Hello World

	// var name type = expression
	var number int  = 2

	// name := expression (short declaration)
	boolean := true

	fmt.Println(number) // 2
	fmt.Println(boolean) // true

	// *** POINTERS ***

	// x := 1
	var x int = 1 

	// p := &x
	var p *int = &x
	
	fmt.Println(p) // 0x140000ac008
	fmt.Println(*p) // 1 (content of pointer)

	// *** TYPE ***

	// fahrenheit & celsius

	type fahrenheit int
	type celsius int
	
	var f fahrenheit = 32
	var c celsius = 0

	c = celsius((f - 32) * 5 / 9)

	fmt.Println(f,c) // 32 0
}