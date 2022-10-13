package main

import (
	"fmt"
)

func greet(c chan string) {
	// listen to value being pushed into channel
	name := <-c
	fmt.Println("Hello", name)
}

func greetUntilQuit(c chan string, quit chan int) {
	for {
		select {
		case name := <-c:
			fmt.Println("Hello", name)
		case <-quit:
			fmt.Println("quitting greeter")
			return
		}
	}
}

func nameReceiver(c chan string, quit chan int) {
	for {
		name, more := <-c
		if more {
			fmt.Printf("A received name: %s\n", name)
		} else {
			fmt.Println("Received all data")
			quit <- 0
		}
	}
}

func nameProducer(c chan string) {
	c <- "Banana"
	c <- "Apple"
	c <- "Orange"
}

func main() {
	// create a new channel
	c := make(chan string)

	// run the channel in a goroutine
	go greet(c)

	// send value to channel
	c <- "World"

	quit := make(chan int)
	go greetUntilQuit(c, quit)
	// you can send multiple data into channels
	// making it like a messaging queue
	c <- "Banana"
	c <- "Apple"
	c <- "Orange"
	quit <- 0

	go nameReceiver(c, quit)
	nameProducer(c)

	// closed channel trigger a signal too
	// it is okay if you don't close a channel
	// it will be garbage collected by golang eventually
	close(c)
	<-quit
}
