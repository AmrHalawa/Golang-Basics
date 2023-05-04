package main

import (
	"fmt"
	"strconv"
	"strings"
)


func main() {
	s := "Hello World"

	fmt.Println(s) // Hello World
	fmt.Println(len(s)) // 11

	fmt.Println(s[0]) // shows the ascii value of that character --> 72
	fmt.Printf("%c", s[0]) // shows the character --> H
	fmt.Println()

	fmt.Println(s[0:5]) // accessing a certain portion // Hello
	fmt.Println(s[:5]) // without providing the start index, will auto start from zero // Hello
	fmt.Println(s[6:]) // World

	// string concatenation
	s = s + " Again"
	s += " Again"
	fmt.Println(s) // Hello World Again Again

	// string literals
	fmt.Println("Hello \nWorld") // Hello (newLine) World
	fmt.Println("Hello \tWorld") // Hello	World (tab)
	fmt.Println("Hello \bWorld") // HelloWorld (backspace)

	// strings are slice of bytes
	// Slice --> is a data structure in go, same as ArrayList in java
	abc := "abc" // string
	b := []byte(abc) // byte slice
	fmt.Println(abc, b) // abc [97 98 99]
	fmt.Printf("%s, %s", abc, b) // abc abc
	fmt.Println()

	abc2 := string(b)
	fmt.Println(abc2) // abc

	// string library
	fmt.Println(strings.ToUpper("Hello World")) // HELLO WORLD
	fmt.Println(strings.ToLower("Hello World")) // hello world
	fmt.Println(strings.HasPrefix("Hello World","Hello")) // true
	fmt.Println(strings.HasSuffix("Hello World", "Hello")) // false
	fmt.Println(strings.HasSuffix("Hello World", "World")) // true
	fmt.Println(strings.Replace("Hello World World", "World", "Hello", 1)) // Hello Hello World
	fmt.Println(strings.Replace("Hello World World", "World", "Hello", 2)) // Hello Hello Hello

	// string builders --> a tool that you can utilize to build your strings
	var sb strings.Builder
	fmt.Println("This is a String Builder:", sb.String()) // empty
	fmt.Println(sb.Cap()) // 0 --> no need for the string builder to allocate any space yet
	fmt.Println(sb.Len()) // 0 --> empty string lives in the string builder
	sb.WriteString("Hello")
	fmt.Println("This is a String Builder:", sb.String())
	fmt.Println(sb.Cap()) // 0 2 4 8 --> and keep on doubling as we find out that we don't have enough space for the string builder
	fmt.Println(sb.Len()) // 5
	sb.Grow(100) // we can increase it manually, by adding more capacity
	fmt.Println(sb.Cap()) // 116 --> 2*cap(b.buf)+n)
	fmt.Println(sb.Len()) // 5
	fmt.Println(sb.String()) // Hello
	sb.Reset()
	fmt.Println(sb.Cap()) // 0 
	fmt.Println(sb.Len()) // 0
	fmt.Println(sb.String()) // empty again

	sb.Write([]byte{65,65,65})
	fmt.Println(sb.String()) // AAA

	// strings & other data types conversion
	x := 123
	y := strconv.Itoa(x)
	fmt.Println(y) // 123
	fmt.Printf("%T", y) // string
	fmt.Println()
	z, _ := strconv.Atoi(y) // convert string back to integer // _ --> unique way in handling errors in golang cause the Atoi() is returning (int, error)
	fmt.Println(z) // 123
	fmt.Printf("%T", z) // int
	println()
}