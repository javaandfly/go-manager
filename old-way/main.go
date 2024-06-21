package main

import (
	"fmt"
	"sync"
)

// 多携程
func main() {

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("i like go 1")
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println("i like go 2")
		}()
	}()

	wg.Wait()
}
