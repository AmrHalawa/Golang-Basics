package main

import (
	"fmt"
	"math"
	"math/big"
	"time"
)

func main () {
	fmt.Println("Hello world 1")

	// declaration of constant
	var five int = 5
	fmt.Println(five)
	var six int = 6
	six = 10
	fmt.Println(six)

	const pi = 3.14
	fmt.Println(math.Pi)

	//multiple constants
	const (
		a = 1
		b
		c = 3
		d 
	)
	fmt.Println(a, b, c, d) // 1 1 3 3

	const (
		zero int = iota + 5
		one
		two
	)
	fmt.Println(zero, one, two) // 5 6 7

	const (
		_ = 1 << (10 * iota) // binary shift to the left 10 times
		kb
		mb
		gb
		tb
		pb
		eb
		yb
	)
	fmt.Println(kb, mb, gb) // 1024 1048576 1073741824 // number of bytes
	// fmt.Println(yb) // untyped constant (overflow) --> much larger than what an integer type can hold
	fmt.Println(yb/eb) // 1024

	// constant.Value()

	fmt.Println(math.Pow(2,100))

	// const largeNum = math.Pow(2,100) // this is not a constant, this method is not returning constant, because amount of power is getting calculated during runtime --> going to return a variable instead of constant which a value is getting assigned during compile time

	// built-in common constants
	fmt.Println(math.Pi) // 3.141592653589793
	fmt.Println(time.February) // February
	fmt.Println(time.Second) // 1s
	fmt.Println(time.UTC) // time zone // UTC
	fmt.Println(big.MaxExp) // 2147483647
	fmt.Println(big.MinExp) // -2147483648
}