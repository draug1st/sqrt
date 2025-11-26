package main

import "fmt"

type Worker interface {
	Work()
}

type WorkerImpl struct {
	WorkDone bool
}

func (w *WorkerImpl) Work() {
	fmt.Println("Working...")
	w.WorkDone = true
}

func NewWorker() Worker {
	var w *WorkerImpl = nil
	return w
}

func main() {
	worker := &WorkerImpl{}
	worker.Work()
	fmt.Println("Work done:", worker.WorkDone)

	var w Worker = NewWorker()
	if w == nil {
		fmt.Println("Worker is nil")
	} else {
		fmt.Println("Worker is not nil")
	}
}
