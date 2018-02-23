# GOLANG-WORKER-QUEUE

Golang worker queue is a simple golang concurrent task runner.

## Installation
```
go get github.com/pgbrtecnologia/golang-worker-queue
```

## Usage
```
package main

import (
	"fmt"
	"time"

	queue "github.com/pgbrtecnologia/golang-worker-queue"
)

type Job struct {
	ID int
}

func (j Job) Run() {
	fmt.Printf("%d started\n", j.ID)
	time.Sleep(200 * time.Millisecond)
	fmt.Printf("%d finished\n", j.ID)
}

func main() {
	dispatcher := queue.GetDispatcher()

	dispatcher.SetWorkerCount(2)

	job1 := Job{ID: 1}
	job2 := Job{ID: 2}
	job3 := Job{ID: 3}

	dispatcher.AddJob(job1)
	dispatcher.AddJob(job2)
	dispatcher.AddJob(job3)

	time.Sleep(500 * time.Millisecond)

	dispatcher.StopAllWorkers()
}
```