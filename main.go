package main

import (
	"fmt"
	"time"
)

func expensiveFunction(key string) string {
	time.Sleep(5 * time.Second)
	return fmt.Sprint("data-", key, ":success")
}

type cache struct {
	storage map[string]string
}

func (c *cache) set(key, value string) {
	c.storage[key] = value
}

func (c *cache) get(key string) string {
	v, ok := c.storage[key]
	if !ok {
		return ""
	}
	return v
}

func main() {
	key := "mydata1"
	start := time.Now()
	result := expensiveFunction(key)
	fmt.Println("expensive function called, duration:", time.Since(start))
	fmt.Println(result)

	// cache the result
	cacher := &cache{
		storage: map[string]string{},
	}
	cacher.set(key, result)
	start = time.Now()

	newResult := cacher.get(key)
	if newResult == "" {
		newResult = expensiveFunction(key)
	}

	fmt.Println("expensive function called, duration:", time.Since(start))
	fmt.Println(newResult)

}
