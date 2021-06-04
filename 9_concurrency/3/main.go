package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	counter := 0
	var wg sync.WaitGroup
	wg.Add(100)
	for i:=0; i<100; i++{
		go func() {
			localVar := counter
			runtime.Gosched()
			localVar++
			counter = localVar
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(counter)
}
