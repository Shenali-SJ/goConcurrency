package main

import (
	"fmt"
	"sync"
)

var x = 0

func increment(wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	x = x + 1
	m.Unlock()
	wg.Done()
}

func main() {
	var wGroup sync.WaitGroup
	var mutex sync.Mutex
	for i := 0; i < 1000; i++ {
		wGroup.Add(1)
		go increment(&wGroup, &mutex)
	}
	wGroup.Wait()
	fmt.Println("Final value of x ", x)
}
