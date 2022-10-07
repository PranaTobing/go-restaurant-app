package main

import (
	"errors"
	"fmt"
	"math/cmplx"
)

func main() {
	// we can declare variables two ways in Golang:

	// long declaration
	var variableName1 string = "hello world!"
	// short declaration
	variableName2 := "hello world!"

	// note: we generally use camel case in Golang
	// as defined in https://go.dev/doc/effective_go#mixed-caps

	// read more about variable declarations here
	// https://go.dev/ref/spec#Variable_declarations

	fmt.Println(variableName1)
	fmt.Println(variableName2)

	// data types in Golang
	// boolean
	boolvar := true
	fmt.Printf("Type: %T Value: %v\n", boolvar, boolvar)

	// string
	stringVar := "hello world!"
	fmt.Printf("Type: %T Value: %v\n", stringVar, stringVar)

	// integer
	// consists of int  int8  int16  int32  int64
	// note: int size depends on system architecture
	intVar1 := int32(1)
	intVar2 := int64(2)
	fmt.Printf("Type: %T Value: %v\n", intVar1, intVar1)
	fmt.Printf("Type: %T Value: %v\n", intVar2, intVar2)

	// unsigned integer
	// consists of uint uint8 uint16 uint32 uint64 uintptr
	// note: uint and uintptr size depends on system architecture
	uintVar1 := uint32(1<<32 - 1)
	uintVar2 := uint64(1<<64 - 1)
	fmt.Printf("Type: %T Value: %v\n", uintVar1, uintVar1)
	fmt.Printf("Type: %T Value: %v\n", uintVar2, uintVar2)

	// float
	floatVar1 := float32(3.5)
	floatVar2 := float64(4.2)
	fmt.Printf("Type: %T Value: %v\n", floatVar1, floatVar1)
	fmt.Printf("Type: %T Value: %v\n", floatVar2, floatVar2)

	// byte > alias for uint8
	byteVarInt := byte(5)
	fmt.Printf("Type: %T Value: %v\n", byteVarInt, byteVarInt)

	byteVarString := []byte("hello world!")
	fmt.Printf("Type: %T Value: %v\n", byteVarString, byteVarString)
	// byte can be casted into string
	fmt.Printf("Type: %T Value: %v\n", byteVarString, string(byteVarString))

	// rune > alias for int32, represents a unicode code point
	runeVar := '😀'
	// notice how rune var is initiated by single quotes, while string uses double quotes
	// single/double quotes are not interchangable in Golang.
	fmt.Printf("Type: %T Value: %v\n", runeVar, runeVar)

	// complex
	// consists of complex64 complex128
	complexVar := cmplx.Sqrt(-7 + 3i)
	fmt.Printf("Type: %T Value: %v\n", complexVar, complexVar)

	// honorable mentions
	var myInterface interface{}
	myInterface = 5
	fmt.Printf("Type: %T Value: %v\n", myInterface, myInterface)
	myInterface = "hello"
	fmt.Printf("Type: %T Value: %v\n", myInterface, myInterface)

	var errVar error = errors.New("an error")
	fmt.Printf("Type: %T Value: %v\n", errVar, errVar)
}
