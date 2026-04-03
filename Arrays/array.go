package main

import "fmt"

func main(){
	// 1. Declare an array of integers with a length of 5 and assign values to it
	var numbers [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println(numbers)

	// 2. Declare an array of strings with a length of 3 and assign values to it
	var fruits [3]string = [3]string{"Apple", "Banana", "Cherry"}
	fmt.Println(fruits)

	// 3. Declare an array of integers without assigning values to it
	var scores [5]int
	fmt.Println(scores) // output -> [0 0 0 0 0] (default values for int)

	// 4. Declare an array of strings without assigning values to it
	var names [3]string
	fmt.Println(names) // output -> ["" "" ""] (default values for string)

	// 5. Declare an array of floats without assigning values to it
	var heights [4]float64
	fmt.Println(heights) // output -> [0 0 0 0] (default values for float64)

	// 6. Declare an array of integers using shorthand syntax
	numbers2 := [5]int{10, 20, 30, 40, 50}
	fmt.Println(numbers2)

	// 7. Declare an array of strings using shorthand syntax
	fruits2 := [3]string{"Mango", "Pineapple", "Grapes"}
	fmt.Println(fruits2)

	fmt.Printf("Length of numbers array: %d\n", len(numbers)) // output -> Length of numbers array: 5
	fmt.Println("First Element of the numbers array is : ",numbers[0]) // output -> First Element of the numbers array is :  1

  // 8. Creating a loop to take input from the user and store it in an array
	var userInput [3]string
	for i := range len(userInput) {
		fmt.Printf("Enter a string for index %d: ", i)
		fmt.Scanln(&userInput[i])
	}
	fmt.Println("User Input:")
	for i := range userInput {
		fmt.Printf("Index %d: %s\n", i, userInput[i])
	}

	// 9. Short hand syntax for printing the elements of an array
	for i, fruit := range fruits {
		fmt.Printf("Index %d: %s\n", i, fruit)
	}

  // 10. 2D Array
	matrix := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println("2D Array (Matrix):", matrix)	


	// Advantages of arrays:
	// 1. Fixed Size: Arrays have a fixed size, which can help prevent memory leaks and improve performance.
	// 2. Contiguous Memory: Arrays store elements in contiguous memory locations, which allows for efficient access and iteration.
	// 3. Type Safety: Arrays are type-safe, meaning that all elements must be of the same type, which can help prevent errors and improve code readability.
	// 4. Built-in Functions: Go provides built-in functions for working with arrays, such as len() for getting the length of an array and copy() for copying elements from one array to another.

	// Built-in functions for arrays in Go:
	// 1. len(): Returns the length of an array.
	// 2. copy(): Copies elements from one array to another.
	// 3. append(): Although append() is primarily used for slices, it can also be used to create a new array by appending elements to an existing array.
	// 4. range: Used in for loops to iterate over the elements of an array.

	// 11. Copying an array using the copy() function
	source := [5]int{1, 2, 3, 4, 5}
	destination := [5]int{}
	copy(destination[:], source[:]) //  We use slicing ([:]) to convert the arrays to slices, as copy() works with slices.
	fmt.Println("Source array:", source)
	fmt.Println("Destination array:", destination)

	// 12. Appending to an array using the append() function (creates a new array)
	original := [3]int{1, 2, 3}
	newArray := append(original[:], 4, 5) //  We use slicing ([:]) to convert the array to a slice, as append() works with slices.
	fmt.Println("Original array:", original)
	fmt.Println("New array after appending:", newArray)

	// Difference between arrays and slices in Go:
	// 1. Size: Arrays have a fixed size defined at compile time, while slices are dynamic and can grow or shrink in size.
	// 2. Memory: Arrays are stored in contiguous memory locations, while slices are references to an underlying array and can point to different parts of the array.
	// 3. Functionality: Slices provide more functionality than arrays, such as the ability to append elements, create sub-slices, and use built-in functions like len() and cap().
	// 4. Performance: Arrays can be more efficient in terms of memory and performance for small, fixed-size collections, while slices are more flexible and suitable for larger, dynamic collections.

}

