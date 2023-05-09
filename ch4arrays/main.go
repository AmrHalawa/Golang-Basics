package main

import "fmt"

func main () {

	var a [3]int
	fmt.Println(a) // [0 0 0] // print the array elements
	fmt.Println(a[0]) // access individual elements by given index // 0
	fmt.Println(len(a)) // 3 // the number of elements in the array
	fmt.Println()
	
	a[0] = 1
	fmt.Println(a) // [1 0 0]

	// in golang --> arrays are values instead of addresses
	aCopy := a
	fmt.Println(aCopy == a) // true
	fmt.Println(aCopy) // [1 0 0]
	a[0] = 2
	fmt.Println(a) // [2 0 0]
	fmt.Println(aCopy) // [1 0 0]
	// when we're declaring the variable aCopy, we're not getting a copy of the address but we're getting a copy of the elements, so they are no longer connected
	fmt.Println(aCopy == a) // false

	b := [4]int{1, 2, 3, 4}
	fmt.Println(b)

	c := [...]int{1, 2}// will cerate an array of 2 elements, if we don't want 4 elements 
	fmt.Println(c) // [1, 2]

	d := [3]int{1, 2}
	fmt.Println(d) // [1, 2, 0]
	// d = [4]int{} // cannot reassign an array of different size to an array of certain size previously
	fmt.Println(a == d) // false // have different elements
	
	for _, v := range b {
		fmt.Println(v) // 1 2 3 4
	}
	fmt.Println()

	array := [...]int{0:1}
	fmt.Println(array) // [1]
	array2 := [...]int{1:1, 4:5}
	fmt.Println(array2) // [0, 1, 0, 0, 5]

	strArray := [...]string{"string1", "string2"}
	fmt.Println(strArray) // [string1 string2]

	array2d := [2][2]int{{1, 2}, {3, 4}}
	fmt.Println(array2d) // [[1 2] [3 4]]

	array3d := [2][2][2]int{array2d, array2d}
	fmt.Println(array3d) // [[[1 2] [3 4]] [[1 2] [3 4]]]
}