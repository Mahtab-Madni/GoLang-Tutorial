package main

import "fmt"
func main() {  
	// Difference between arrays and slices:
	// 1. Size: Arrays have a fixed size that is determined at compile time, while slices are dynamic and can grow or shrink in size as needed.
	// 2. Memory: Arrays are stored in contiguous memory locations, while slices are references to an underlying array and can point to different parts of the array.
	// 3. Functionality: Slices provide more functionality than arrays, such as the ability to append elements, create sub-slices, and use built-in functions like len() and cap().
	// 4. Performance: Arrays can be more efficient in terms of memory and performance for small, fixed-size collections, while slices are more flexible and suitable for larger, dynamic collections.

	// 1. Declare a slice of integers and assign values to it
	var numbers []int = []int{1, 2, 3, 4, 5}
	fmt.Println(numbers)

	// 2. Short hand syntax for declaring a slice of strings and assigning values to it
	fruits := []string{"Apple", "Banana", "Cherry"}
	fmt.Println(fruits)

	// 3. Declare a slice of integers without assigning values to it
	var scores []int  // store nil value in scores variable (nil == null == No value)
	fmt.Println(scores) // output -> [] (default value for a slice is nil, which is represented as an empty slice when printed)

	// Built-in functions for slices in Go:
	// 1. len(): Returns the length of a slice. The length is the number of elements currently in the slice.
	// 2. cap(): Returns the capacity of a slice. The capacity is the total number of elements that the slice can hold before it needs to be resized. When a slice is created, it has an underlying array that can hold a certain number of elements. If the slice exceeds its capacity, Go will automatically create a new underlying array with a larger capacity and copy the existing elements to it.
	// 3. append(): Adds elements to the end of a slice and returns a new slice with the added elements.
	// 4. copy(): Copies elements from one slice to another.
	// 5. range: Used in for loops to iterate over the elements of a slice.

	// 4. Using the append() function to add elements to a slice
	numbers = append(numbers, 6, 7, 8)
	fmt.Println(numbers) // output -> [1 2 3 4 5 6 7 8]

	// 5. Using the copy() function to copy elements from one slice to another
	source := []int{1, 2, 3, 4, 5}
	destination := make([]int, len(source)) // We create a new slice called destination with the same length as the source slice using the make() function.
	copy(destination, source)
	fmt.Println("Source slice:", source)
	fmt.Println("Destination slice:", destination)

	// 6. Creating a sub-slice from an existing slice
	subSlice := numbers[2:5] // This creates a new slice called subSlice that contains the elements from index 2 to index 4 of the numbers slice (the end index is exclusive).
	fmt.Println(subSlice) // output -> [3 4 5]

	// 7. Using the range keyword to iterate over a slice
	fmt.Println("Iterating over numbers slice:")
	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	} // output -> Index: 0, Value: 1

	// 8. Using the range keyword to iterate over a slice without using the index
	fmt.Println("Iterating over numbers slice without index:")
	for _, value := range numbers {
		fmt.Printf("Value: %d\n", value) // output -> Value: 1
	}

  // 9. creating a loop to take input from the user and store it in a slice
	var userInput []string	
	for i := range 3 { 
		var input string
		fmt.Printf("Enter a string for index %d: ", i)
		fmt.Scanln(&input)
		userInput = append(userInput, input) // We use the append() function to add the user input to the userInput slice.
	}
	fmt.Println("User input:", userInput)
	
}