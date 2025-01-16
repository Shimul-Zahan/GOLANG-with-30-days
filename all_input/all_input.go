package allinput

import (
	"bufio"
	"fmt"
	"os"
)

func Inputs() {
	// For a single-word input
	var name string
	fmt.Println("Enter your name (single word): ")
	// fmt.Scan(&name) // Scans a single word (until a space)
	fmt.Println("Your name is:", name)

	// For multi-word input
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your full name (multi-word): ")
	// The bufio.ReadString function returns two values:
	// - The string that was read until the newline ('\n')
	// - An error, if any occurred during reading
	// We use "_" to discard the error value here since we are not handling it
	fullName, _ := reader.ReadString('\n')
	fmt.Println("Your full name is:", fullName)
}
