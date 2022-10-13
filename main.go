package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func testPrint(id int) {
	for i := 0; i < 10; i++ {
		fmt.Printf("[%d] counting %d\n", id, i)
		amt := time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond * amt)
	}
}

func main() {
	go testPrint(0)

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			testPrint(i)
		}(i)
	}

	// what happens when you remove this wg wait?
	wg.Wait()

	var sharedResource string
	mutex := &sync.Mutex{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			// try doing it without mutex and see what happens!
			mutex.Lock()
			fmt.Println("Previous value: ", sharedResource)
			sharedResource = fmt.Sprintf("key owned by: [%d]", id)
			fmt.Println(sharedResource)
			mutex.Unlock()
		}(i)
	}
	wg.Wait()

	// this is random
	fmt.Println("final resource: ", sharedResource)
}
