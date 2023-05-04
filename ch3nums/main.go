package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Hello")

	// integers

	var x = 1 // int32 int64
	var x32 int32 = 1
	var x64 int64 = 1
	fmt.Println(x) // 1
	fmt.Println(x32) // 1
	fmt.Println(x64) // 1

	fmt.Printf("%T, %T, %T", x , x32, x64) // int, int32, int64
	fmt.Println()

	x = int(x32)
	fmt.Println(x) // 1
	x = int(x64)
	fmt.Println(x) // 1

	var unsigned uint = 1
	fmt.Println(unsigned)

	y := 2
	fmt.Println(x + y) // 3
	fmt.Println(x % y) // 1
	fmt.Println(x == y) // false
	fmt.Println(x < y) // true

	// floating point numbers

	pi := 3.141
	fmt.Println(pi)
	fmt.Printf("%T", pi)
	fmt.Println()

	var pi32 float32 = 3.141
	fmt.Printf("%T",pi32)
	fmt.Println()

	pi = float64(3.141)
	fmt.Println(pi)

	number := 1e100
	fmt.Println(number) // 1e+100

	// integer & floating point conversion

	a := 1
	b := 3.141
	fmt.Println(a) // 1
	fmt.Println(b) // 3.141


	b = float64(a)
	fmt.Println(b) // 1.00000

	a = int(b)
	fmt.Println(a) // 3

	// Math Library

	fmt.Println("Ceiling", math.Ceil(3.000001)) // 4
	fmt.Println("Floor", math.Floor(3.000001)) // 3
	fmt.Println("Min", math.Min(1,0)) // 0
	fmt.Println("Max", math.Max(1,0)) // 1
	fmt.Println("Absolute", math.Abs(-1)) // 1
	fmt.Println("Square" ,math.Sqrt(100)) // 10
	fmt.Println("Power" ,math.Pow(2,3)) // 8

	// ** bonus ** 
	// complex numbers
	fmt.Println("Complex" ,complex(1,2)) // (1+2i)
}