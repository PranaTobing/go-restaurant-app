package main

import (
	"fmt"
)

func print_hello() {
	fmt.Printf("Hello!\n")
}

func main() {
	// we'll see the effects later
	// defer runs the function last, after this function finishes
	defer print_hello()

	age := 5

	// in if conditionals,
	// there are if, else if, and else
	if age < 18 {
		fmt.Printf("Not old enough!\n")
	} else if age > 60 {
		fmt.Printf("Too old!\n")
	} else {
		fmt.Printf("Welcome!\n")
	}

	// if have too many conditionals, you can try using a switch
	grade := "B"

	switch grade {
	case "A":
		fmt.Printf("Good job!\n")
	case "B", "C":
		fmt.Printf("Nice work!\n")
		fmt.Printf("You can do better!\n")
	case "F":
		fmt.Printf("Please see the teacher\n")
		// in golang, there is no break statement,
		// instead, you have to use fallthrough to indicate
		// you are making this code go to the next case
		fallthrough
	default:
		fmt.Printf("You have to try again!\n")
	}

	// there are many ways to run a loop in Golang

	// the first one is the three-component loop
	// it consists of a statement, a condition, and a post statement
	for i := 1; i < 5; i++ {
		fmt.Printf("Looping %d\n", i)
	}

	// the second one is a while loop
	j := 0
	for j < 5 {
		fmt.Printf("Looping %d\n", j)
		j += 1
	}

	i := 0
	// infinite loop
	for {
		if i == 3 {
			break
		}
		fmt.Printf("I will keep repeating unless stopped\n")
		i += 1
	}

	numbers := []int64{1, 2, 3, 4, 5}
	for i, v := range numbers {
		fmt.Printf("I got %d on index %d\n", v, i)
	}
}
