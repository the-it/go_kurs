package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(100)
	counter := 0
	var mutex sync.Mutex

	for i:=0; i<100; i++{
		go func() {
			mutex.Lock()
			localVar := counter
			localVar++
			counter = localVar
			mutex.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(counter)
}
