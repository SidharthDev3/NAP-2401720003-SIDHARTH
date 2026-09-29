package channels

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker %d started Job %d \n", id, job)
		time.Sleep(time.Millisecond * 500)
		results <- job * 2
	}
}

func ChannelsQuestion() {
	//define numJobs
	numJobs := 100

	//create jobs and results channels
	jobs := make(chan int, numJobs)
	result := make(chan int, numJobs)

	//start workers
	for w := 1; w <= numJobs; w++ {
		go worker(w, jobs, result)
	}

	//send jobs
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}

	//close channel
	close(jobs)

	//collect results from receiver channels
	for a := 1; a <= numJobs; a++ {
		fmt.Println("Result: ", <-result)
	}
}
