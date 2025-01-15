package prints

import (
	"fmt"
	"os"
)

func PrintsFunctions() {
	// Using fmt.Print
	fmt.Print("Hello, ")
	fmt.Print("World!")

	// Using fmt.Println
	fmt.Println("\nHello, World!") // Adds a newline after the output

	// Using fmt.Printf
	name := "Shimul"
	age := 25
	fmt.Printf("Name: %s, Age: %d\n", name, age)

	// Using fmt.Sprintf (returns formatted string)
	result := fmt.Sprintf("Name: %s, Age: %d", name, age)
	fmt.Println(result)

	// Using fmt.Fprint (write to an output stream)
	fmt.Fprint(os.Stdout, "Printing to standard output stream.\n")

	// Using fmt.Fprintf (formatted output to a writer)
	fmt.Fprintf(os.Stdout, "Formatted Age: %d, Weight: %.2f\n", age, 70.5)

	// Using fmt.Sprintln (adds newline)
	str := fmt.Sprintln("This string will end with a newline.")
	fmt.Print(str)
}
