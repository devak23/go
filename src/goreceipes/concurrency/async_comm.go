package concurrency

import (
	"time"
	"goreceipes/concurrency/syncutils"
	"fmt"
	"math/rand"
)

// define an interface Runnable that has a run method
type Runnable interface {
	Run()
}

// define a Task struct that holds data for a Task
type Task struct {
	Id int
	JobId int
	Status string
	CreatedOn time.Time
}

// The main program
func AsyncCommMain() {
	// signal the wait group that we have 3 worker threads
	syncutils.Wg.Add(3)

	// create a buffered queue(channel) of 10 tasks
	queue := make (chan *Task, 10)

	// Launch the goroutines that handle the work
	for i := 0; i < 3; i++ {
		go worker(queue, i)
	}

	// fill up the queue with Tasks
	for i:=0; i<10; i++ {
		queue <- &Task{
			Id: i,
			JobId: 100 + i,
			CreatedOn: time.Now(),
		}
	}

	// close the queue
	close(queue)

	// Wait for all the work to be completed
	syncutils.Wg.Wait()
}

// worker is launched as goroutine processes task from the queue
// The taskChannel is marked as readonly.
func worker(taskChannel <- chan *Task, workerId int) {
	defer syncutils.Wg.Done()

	// Run the task for each entry from the channel
	for task := range taskChannel {
		fmt.Printf("Worker-%d: received request for TaskId: %d - Job: %d\n", workerId, task.Id, task.JobId)

		// Execute the task
		task.Run()

		// display the finished work
		fmt.Printf("Worker-%d: completed request for TaskId: %d - Job: %d, Status = %s\n", workerId, task.Id, task.JobId, task.Status)
	}
}

func (t *Task) Run() {
	// introducing a delay to simulate working
	sleepDuration := rand.Int63n(1000)
	time.Sleep(time.Duration(sleepDuration)  * time.Millisecond)
	t.Status = "Completed"
}
