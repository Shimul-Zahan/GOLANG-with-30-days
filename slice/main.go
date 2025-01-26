package main

import "fmt"

var slices = []string{
	"Robotix exam",
	"Data Science Projects",
	"Golang Learning",
	"Database Learning with Projects",
}

var arr [100]string

func main() {
	var name string = "shimul"
	fmt.Printf("%s \n", name)

	copy(arr[:], slices)
	fmt.Println("Array after adding slices:", arr)

	// View the array elements (skip empty values)
	println("--------view array elements----------")
	viewArray()

	// Display the slice elements
	println("--------view slices----------")
	viewList()

	// Add a new task and display the updated slice
	addTask("Go and Explore all")

	println("--------view updated slices----------")
	viewArray()
}

func viewList() {
	for index, slice := range slices {
		fmt.Printf("%d. %s\n", index+1, slice)
	}
}

func addTask(task string) {
	slices = append(slices, task)
}

func viewArray() {
	for i := 0; i < len(slices); i++ {
		if arr[i] != "" {
			fmt.Println(arr[i])
		}
	}
}
