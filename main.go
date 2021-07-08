package main

import (
	"fmt"
	"sync"
	"time"
)

func hello(channel chan bool) {
	fmt.Println("Hello world goroutine")
	channel <- true   //writing to the channel
}

func calcSquare(num int, square chan int) {
	sum := 0
	for num != 0 {
		digit := num % 10
		sum += digit * digit
		num /= 10
	}

	square <- sum  //writing sum value to the channel
}

func calcCubes(num int, cube chan int) {
	sum := 0
	for num != 0 {
		digit := num % 10
		sum += digit * digit * digit
		num /= 10
	}
	cube <- sum
}

func sendData(channel chan <- int) {
	channel <- 19
}

func printNum(channel chan int) {
	for i := 0; i < 10; i++ {
		channel <- i
	}
	close(channel)
}

func writeNum(channel chan int) {
	for i := 0; i < 5; i++ {
		channel <- i
		fmt.Println("Wrote ", i, " to channel")
	}
	close(channel)
}

//pointer should be passed
//else each goroutine will have their own copy of waitGroup
//bc functions are pass by value
func process(i int, waitGroup *sync.WaitGroup) {
	fmt.Println("Started goroutine ", i)
	time.Sleep(2 * time.Second)
	fmt.Println("Ended goroutine ", i)
	waitGroup.Done()  //decrement counter
}

func main() {
	done := make(chan bool)   //creating a channel
	go hello(done)

	//reading from the channel
	//data received is not stored - legal
	//main go routine is blocked until data is received from the channel
	<- done

	//time.Sleep(1 * time.Second)
	fmt.Println("main function - after hello()")

	//2nd example
	number := 678
	square := make(chan int)
	cube := make(chan int)

	go calcSquare(number, square)
	go calcCubes(number, cube)

	//reading from channels
	squares, cubes := <-square, <-cube
	fmt.Println("Final output ", squares + cubes)

	//example 3
	channel := make(chan int)
	go sendData(channel)
	fmt.Println(<-channel)

	//example 4
	ch := make(chan int)
	go printNum(ch)
	for {
		v, ok := <- ch
		if ok == false {
			break
		}
		fmt.Println("received ", v, ok)
	}

	fmt.Println()

	//example 5
	ch2 := make(chan int)
	go printNum(ch2)
	for v := range ch2 {
		fmt.Println("Received ", v)
	}

	fmt.Println()

	//example 6 - buffered channels
	channel2 := make(chan string, 2)

	channel2 <- "Shenali"
	channel2 <- "Jayakody"

	fmt.Println(<-channel2)
	fmt.Println(<-channel2)

	fmt.Println()

	//example 7
	channel3 := make(chan int, 2)
	go writeNum(channel3)
	time.Sleep(2 * time.Second)
	for v := range channel3 {
		fmt.Println("Read ", v, " from channel3")
		time.Sleep(2 * time.Second)
	}

	fmt.Println()

	//waitGroup
	no := 3
	var wGroup sync.WaitGroup   //zero value waitGroup is created.
	for i := 0; i < no; i++ {
		wGroup.Add(1)  //increment counter of waitGroup by one
		go process(i, &wGroup)  //goroutine
	}

	wGroup.Wait()  //blocks the main goroutine until the counter of the waitGroup becomes zero
	fmt.Println("All go routines finished executing")

}
