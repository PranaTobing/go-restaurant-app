package main

import (
	"errors"
	"fmt"
)

// in golang, functions are declared like this
// you have:
// the function name (add)
// the input variables, also with their expected types (var1,var2)
// and the expected output variable of this function (an integer)
func add(var1 int, var2 int) int {
	return var1 + var2
}

// you can also declare input types in one sequence like so
// all of the variables declared before the type will use that type
func subtract(var1, var2 int) int {
	return var1 - var2
}

// you can also declare expected output variables for better readability
// this also declares a variable called 'result' internally,
// so you can call it within the function without declaring it
func multiply(var1, var2 int) (result int) {
	result = var1 * var2
	return
}

// what happens when there are errors?
// in Golang, we have another type, called 'error', that can pass our errors through
func divide(var1, var2 int) (float64, error) {
	if var2 == 0 {
		return 0, errors.New("cannot divide by 0")
	}

	// int variables must be typecasted first before processing,
	// or else it will truncate the decimals
	return float64(var1) / float64(var2), nil
}

func main() {
	// main is a function in itself
	// it is not possible to declare a function within a function
	// unless, it is an anonymous function
	anonymousFunction := func() string {
		return "Hello world!"
	}
	fmt.Println(anonymousFunction())

	addResult := add(3, 6)
	fmt.Printf("ran function add, got %d\n", addResult)

	subtractResult := subtract(6, 2)
	fmt.Printf("ran function subtract, got %d\n", subtractResult)

	multiplyResult := multiply(5, 4)
	fmt.Printf("ran function multiply, got %d\n", multiplyResult)

	// to handle multiple outputs, you can use short declaration
	divisionResult, err := divide(5, 3)
	if err != nil {
		fmt.Printf("ran function divide, got error! %s\n", err)
	}
	fmt.Printf("ran function divide, got %f\n", divisionResult)

	// you cannot re-declare variables using short declaration
	// if both of the variables declared already exist before
	divisionResult, err = divide(5, 0)
	if err != nil {
		fmt.Printf("ran function divide, got error! %s\n", err)
	}
	fmt.Printf("ran function divide, got %f\n", divisionResult)
}
