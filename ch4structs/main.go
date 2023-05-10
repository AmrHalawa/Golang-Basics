package main

import "fmt"

func main () {
	type ninja struct {
		// member variables
		name string
		weapons []string
		level int
	}

	// assigning values in struct
	amr := ninja{name: "Amr"}
	amr = ninja{"Amr", []string{"Ninja Star"}, 1}
	fmt.Println(amr) // {Amr [Ninja Star] 1}
	amr.level++
	fmt.Println(amr) // {Amr [Ninja Star] 2}
	amr.weapons = append(amr.weapons, "Ninja Sword")
	fmt.Println(amr) // {Amr [Ninja Star Ninja Sword] 2}

	type dojo struct {
		name string
		ninja ninja
	}

	golangDojo := dojo{
		name: "Golang Dojo",
		ninja: amr, // passed by value
	}
	fmt.Println(golangDojo) // {Golang Dojo {Amr [Ninja Star Ninja Sword] 2}}
	fmt.Println(golangDojo.ninja.level) // 2
	golangDojo.ninja.level = 3
	fmt.Println(golangDojo.ninja.level) // 3
	fmt.Println(amr.level) // 2 // passed by value, when we're referencing amr, when we passing amr variable into the construction of dojo struct, we 're passing it as a copy of wallace variable.
	// when we're modifying ninja object or struct inside of golangDojo struct, we're modifying the copy of amr not amr itself, so it will not reflect

	type addressedDojo struct {
		name string
		ninja *ninja
	}
	addressed := addressedDojo{"Addressed Golang Dojo", &amr}
	fmt.Println(addressed) // {Addressed Golang Dojo 0x1400010c180}
	fmt.Println(*addressed.ninja) // {Amr [Ninja Star Ninja Sword] 2}
	addressed.ninja.level = 4
	fmt.Println(addressed.ninja.level) // 4
	fmt.Println(amr.level) // 4

	// when using the new operator, it's going to return us a pointer instead of the object itself

	aliPointer := new(ninja)
	fmt.Println(aliPointer) // &{ [] 0} // pointer
	fmt.Println(*aliPointer) // { [] 0} // content
	aliPointer.name = "Ali"
	aliPointer.weapons = []string{"Ninja Star"}
	aliPointer.level = 1
	fmt.Println(*aliPointer) // {Ali [Ninja Star] 1}

	intern := ninjaIntern{name: "Intern", level: 1}
	intern.sayHello() // Hello
	intern.sayName() // Intern
}

// adding functions to a struct
type ninjaIntern struct {
	name string
	level int
}

// it's a function that belongs to the struct
func (ninjaIntern) sayHello() {
	fmt.Println("Hello")
}

func (n ninjaIntern) sayName() {
	fmt.Println(n.name)
}

