package main

import (
	"fmt"
)

// 1. Define Animal interface

type Animal interface {
	Speak() string
}

// 2. Define structs Dog, Cat, Cow
type Dog struct{}
type Cat struct{}
type Cow struct{}

// 3. Implement Speak() for each struct
func (d Dog) Speak() string { return "Woof!" }
func (c Cat) Speak() string { return "Meow!" }
func (c Cow) Speak() string { return "Moo!" }

// 4. Function that accepts Animal and prints Speak()

func MakeItSpeak(a Animal) {
	fmt.Println(a.Speak())
}

func main() {
	// 5. Create Dog, Cat, Cow instances
	// 6. Call MakeItSpeak for each
	MakeItSpeak(Dog{})
	MakeItSpeak(Cat{})
	MakeItSpeak(Cow{})
}
