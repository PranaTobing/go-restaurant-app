package main

import (
	"fmt"
)

// pointers are a way for us to manipulate variable values from another place
func incrementPtr(var1 *int) {
	*var1 = *var1 + 1
}

func increment(var1 int) {
	var1 = var1 + 1
}

func main() {
	var1 := 1
	fmt.Printf("my value is %d\n", var1)

	// when you call a function that expects a value,
	// your var1 is not mutated.
	increment(var1)
	fmt.Printf("my value is %d\n", var1)
	// to do so, we can use pointers.

	// you can pass a normal variable as a pointer
	// by sending its address in the memory, marked using &
	incrementPtr(&var1)
	fmt.Printf("my value is %d\n", var1)

	// you can also initialize a variable with pointers
	// by default it will not point to any memory address
	var var2 *int
	fmt.Printf("my value is %v\n", var2)

	// we can call new to manually allocate memory
	var2 = new(int)
	*var2 = 5
	fmt.Printf("my value is %v\n", *var2)
}
