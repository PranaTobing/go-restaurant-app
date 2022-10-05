package main

import (
	"fmt"
	"time"
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

func counter(quit chan int) {
	var i int = 0
	for {
		select {
		case <-quit:
			return
		default:
			fmt.Println(i)
			i += 1
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func timer(quit chan int) {
	for i := 5; i >= 0; i-- {
		time.Sleep(1 * time.Second)
		fmt.Printf("timer %d...\n", i)
	}
	fmt.Println("quitting")
	quit <- 0
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

	go counter(quit)
	go timer(quit)
	<-quit

	// channel should be closed
	close(c)
}
