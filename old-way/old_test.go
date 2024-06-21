package main

import (
	"fmt"
	"sync"
	"testing"
)

func TestMultGoroutine(t *testing.T) {
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
