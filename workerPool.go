package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Job struct {
	id int
	randNum int
}

type Result struct {
	job Job
	sum int
}

var jobs = make(chan Job, 10)
var results = make(chan Result, 10)

//add jobs to the jobs channel
func allocate(noOfJobs int) {
	for i := 0; i < noOfJobs; i++ {
		randomNum := rand.Intn(999)
		//creating the struct
		job := Job{
			id:      i,
			randNum: randomNum,
		}
		//send each job to the jobs channel
		jobs <- job
	}
	close(jobs)  //no more jobs will be sent to the jobs channel
}

func result(done chan bool) {
	for result := range results {
		fmt.Printf("Job id %d, random number %d, sum %d \n", result.job.id, result.job.randNum, result.sum)
	}
	done <- true  //send to the bool channel that the work is done here
}

func createWorkerPool(noOfWorkers int) {
	var wGroup sync.WaitGroup
	for i := 0; i < noOfWorkers; i++ {
		wGroup.Add(1)
		go worker(&wGroup)
	}
	wGroup.Wait()  //wait until the goroutines are executed
	close(results)  //no more data will be sent to the results channel
}

func worker(waitGroup *sync.WaitGroup) {
	for job := range jobs {
		output := Result{job, calculateSum(job.randNum)}
		results <- output  //sending output to the results channel
	}
	waitGroup.Done() //decrement waitGroup
}

func calculateSum(number int) int {
	sum := 0
	no := number
	for no != 0 {
		digit := no % 10
		sum += digit
		no /= 10
	}
	time.Sleep(2 * time.Second)
	return sum
}

func main() {
	startTime := time.Now()
	noOfJobs := 100

	go allocate(noOfJobs)

	done := make(chan bool)
	go result(done)

	noOfWorkers := 20
	createWorkerPool(noOfWorkers)

	<-done

	endTime := time.Now()
	timeTaken := endTime.Sub(startTime)
	fmt.Println("Total time taken ", timeTaken.Seconds(), " seconds.")

}
