package main

import "fmt"

func main () {
	
	// for loops
	 
	for i, j := 0, 1; i < 10 && j < 10; i, j = i + 1, j + 1 {  // this syntax is preferable cause if you won't use i ever again, if you put it inside the for loop, the the computer will know that after the for loop is done, this variable is no longer useful, and just get it out of the stack.
		fmt.Println(i, j)
	}
	fmt.Println()

	// foreach loops --> more modern, preferable way to write write for loops when working with editable data structures

	s := "Hello World"
	for i := 0; i < len(s); i++ {
		fmt.Printf("%d %c", i, s[i]) // print index & character
		fmt.Println()
	}
	fmt.Println()

	fmt.Println("foreach loop")
	for i, c := range s {
		fmt.Printf("%d %c", i, c) // print index & character
		fmt.Println()
	}
	fmt.Println()

	// keywords

	// break
	for _, c := range s {
		if c == ' ' {
			break // will break the for loop
		}
		fmt.Printf("%c", c) // Hello
	}
	fmt.Println()

	// continue
	for _, c := range s {
		if c == ' ' {
			continue // with ignore the current iteration, and ignore everything after the continue keyword, and move on to the next iteration
		}
		fmt.Printf("%c", c) // HelloWorld
	}
	fmt.Println()

	// labels
	iForLoop: // if we want to break the i for loop
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if j == 3 {
				break iForLoop // will not print any j with value of 3
			}
			fmt.Printf("%d%d ", i, j)
		}
		fmt.Println()
	}
}