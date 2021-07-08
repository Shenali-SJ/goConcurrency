package main

import (
	"fmt"
	"sync"
)

var x = 0
var y = 0

func increment(wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	x = x + 1
	m.Unlock()
	wg.Done()
}

func increment2(wg *sync.WaitGroup, channel chan bool) {
	channel <- true
	y = y + 1
	<- channel
	wg.Done()
}

func main() {
	//mutex for race condition
	var wGroup sync.WaitGroup
	var mutex sync.Mutex
	for i := 0; i < 1000; i++ {
		wGroup.Add(1)
		go increment(&wGroup, &mutex)
	}
	wGroup.Wait()
	fmt.Println("Final value of x ", x)

	fmt.Println()

	//channels for race condition
	var wGroup2 sync.WaitGroup
	channel := make(chan bool, 1)
	for i := 0; i < 1000; i++ {
		wGroup2.Add(1)
		go increment2(&wGroup2, channel)
	}
	wGroup2.Wait()
	fmt.Println("Final value of y ", x)
}
