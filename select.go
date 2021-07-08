package main

import (
	"fmt"
	"time"
)

func server1(ch chan string) {
	time.Sleep(2 * time.Second)
	ch <- "from server 1"
}

func server2(ch chan string){
	time.Sleep(3 * time.Second)
	ch <- "from server 2"
}

func main() {
	channel1 := make(chan string)
	channel2 := make(chan string)

	go server1(channel1)
	go server2(channel2)

	select {
	case s1 := <-channel1:
		fmt.Printf(s1)
	case s2 := <-channel2:
		fmt.Printf(s2)
	}
}
