package main

import (
	"fmt"
	"time"
)

type myService struct {
	cacher *cache
}

func (ms *myService) expensiveFunction(key string) string {
	time.Sleep(5 * time.Second)
	return fmt.Sprint("data-", key, ":success")
}

func (ms *myService) GetData(key string) string {
	if ms.cacher != nil {
		if cachedData := ms.cacher.get(key); cachedData != "" {
			return cachedData
		}
	}

	result := ms.expensiveFunction(key)
	ms.cacher.set(key, result)

	return result
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
	// cache the result
	cacher := &cache{
		storage: map[string]string{},
	}
	service := &myService{
		cacher: cacher,
	}

	key := "mydata1"

	start := time.Now()
	fmt.Println("calling expensive function")
	result := service.GetData(key)
	fmt.Println("expensive function called, duration:", time.Since(start))
	fmt.Println(result)

	start = time.Now()
	fmt.Println("calling expensive function")
	newResult := service.GetData(key)
	fmt.Println("cachedexpensive function called, duration:", time.Since(start))
	fmt.Println(newResult)

}
