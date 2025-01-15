package variableanddatatypes

import (
	"fmt"
)

// String variables
var AppName string = "Learning Go"
var Developer string = "Eleush Zahan Shimul"

// Integer variables
var Year int = 2025
var MaxUsers int = 1000

// Float variables
var Pi float64 = 3.14159
var Version float32 = 1.0

// Boolean variable
var IsReleased bool = true

// Complex variable
var ComplexExample complex128 = complex(5, 7)

// Character variable (rune)
var Initial rune = 'L'

// Byte variable
var Symbol byte = '@'

// VariableAndDataTypes demonstrates all data types
func VariableAndDataTypes() {
	fmt.Println("String Variables:")
	fmt.Printf("App Name: %s, Developer: %s\n", AppName, Developer)

	fmt.Println("\nInteger Variables:")
	fmt.Printf("Year: %d, Max Users: %d\n", Year, MaxUsers)

	fmt.Println("\nFloat Variables:")
	fmt.Printf("Pi: %.5f, Version: %.1f\n", Pi, Version)

	fmt.Println("\nBoolean Variable:")
	fmt.Printf("Is Released: %t\n", IsReleased)

	fmt.Println("\nComplex Variable:")
	fmt.Printf("Complex Number: %v\n", ComplexExample)

	fmt.Println("\nCharacter (Rune) Variable:")
	fmt.Printf("Initial: %c\n", Initial)

	fmt.Println("\nByte Variable:")
	fmt.Printf("Symbol: %c\n", Symbol)
}
