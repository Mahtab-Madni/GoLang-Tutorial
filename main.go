package main

import "fmt"

// add is a function that takes two integers as parameters and returns their sum as an integer.
func add(x, y int) int {
	return x + y
}

func main() {
	// 1. Declare a variable of type string and assign it a value
	var greeting string = "Hello, Go!"
	fmt.Println(greeting)

	var msg = "Welcome to Go programming!" // type inference allows us to omit the type when declaring a variable, as Go can infer it from the assigned value.
	fmt.Println(msg)

	// 2. Declare a variable of type int and assign it a value
	var age int = 30
	fmt.Println(age)	

	// 3. Declare a variable of type float64 and assign it a value
	var pi float64 = 3.14
	fmt.Println(pi)

	// 4. Declare a variable of type bool and assign it a value
	var isGoFun bool = true
	fmt.Println(isGoFun)

	// 5. Declare a variable of type string without assigning it a value
	var name string
	fmt.Println(name) // output -> "" (default value for string)

	// 6. Declare a variable of type int without assigning it a value
	var count int
	fmt.Println(count) // output -> 0 (default value for int)

	// 7. Declare a variable of type float64 without assigning it a value
	var temperature float64
	fmt.Println(temperature) // output -> 0 (default value for float64)

	// 8. Declare a variable of type bool without assigning it a value
	var isRaining bool
	fmt.Println(isRaining) // output -> false (default value for bool)

	// 9. Declare a variable of type string and assign it a value using shorthand syntax
	message := "Welcome to Go programming!"
	fmt.Println(message)

	// 10. Declare a variable of type int and assign it a value using shorthand syntax
	score := 100
	fmt.Println(score)

	// 11. Declare a variable of type float64 and assign it a value using shorthand syntax
	radius := 5.5
	fmt.Println(radius)

	// 12. Declare a variable of type bool and assign it a value using shorthand syntax
	isSunny := false
	fmt.Println(isSunny)

	// 13. Declare multiple variables of the same type in a single line
	var x, y, z int = 1, 2, 3
	fmt.Println(x, y, z)

	// 14. Declare multiple variables of different types in a single line
	var a, b, c = "Go", 2024, true
	fmt.Println(a, b, c)

	// 15. Declare a constant of type string and assign it a value
	const language string = "Go"
	fmt.Println(language)

	// 16. Declare a constant of type int and assign it a value
	const year int = 2024
	fmt.Println(year)

	// 17. Declare a constant of type float64 and assign it a value
	const gravity float64 = 9.81
	fmt.Println(gravity)

	// 18. Declare a constant of type bool and assign it a value
	const isProgrammingFun bool = true
	fmt.Println(isProgrammingFun)

	// 19. Declare a for loop that iterates from 1 to 5 and prints the current iteration number
	for i := 1; i <= 5; i++ {
		fmt.Println("Iteration:", i)
	}

	//20. Declare a while loop using a for loop that continues until a variable reaches a certain value
	counter := 0
	for counter < 5 {
		fmt.Println("Counter:", counter)
		counter++
	}

	//21. using range in a loop to iterate over a collection of items
	numbers := []int{1, 2, 3, 4, 5}  // slice of integers
	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value) // output -> Index: 0, Value: 1
	}

	//22. using nested loops to iterate over multiple collections of items
	matrix := [][]int{ // 2D slice (matrix) of integers
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	for i := range matrix { // Outer loop iterates over rows
		for j := range matrix[i] { // Inner loop iterates over columns
			fmt.Printf("Element at [%d][%d]: %d\n", i, j, matrix[i][j])
		}
	}

	// 23. using continue / break in a loop
	for j := 1; j <= 10; j++ {
		if j%2 == 0 {
			continue // Skip even numbers
		}
		if j > 7 {
			break // Exit the loop if j is greater than 7
		}
		fmt.Println("Odd number:", j)
	}

	//24 . using if / else statements to check a condition and execute different code blocks based on the result
	number := 10
	if number%2 == 0 {
		fmt.Println(number, "is an even number.")
	} else {
		fmt.Println(number, "is an odd number.")
	}

	//25. using switch statements to check multiple conditions and execute different code blocks based on the result
	day := "Monday"
	switch day {
		case "Monday":
			fmt.Println("It's the start of the week.")

		case "Friday":
			fmt.Println("It's the end of the week.")

		default:
			fmt.Println("It's a regular day.")
	}
	
	//26. using logical operators (&&, ||, !) to combine multiple conditions in an if statement
	age = 25
	isStudent := true
	if age < 30 && isStudent {
		fmt.Println("You are a young student.")
	}
	
	// 27. Declare a function that takes two integers as parameters and returns their sum
	sum := add(10, 20)
	fmt.Println("Sum:", sum)

	// 28. Declare a function that takes a string as a parameter and prints a greeting message
	greet := func(name string) {
		fmt.Printf("Hello, %s!\n", name)
	}
	greet("Alice")

	// 29. Declare a function that takes no parameters and returns a string
	getMessage := func() string {
		return "Welcome to Go programming!"
	}
	fmt.Println(getMessage())

	// 30. Declare a function that takes a variable number of arguments and returns their average
	average := func(nums ...float64) float64 {
		var sum float64
		for _, num := range nums {
			sum += num
		}
		return sum / float64(len(nums))
	}
	fmt.Println(average(1, 2, 3, 4, 5))
}