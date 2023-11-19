package main

import "sync"

type Worker struct {
	tasks               *[]Task
	results    			*[] Result
	executor     		Executor
	cdl 			    *sync.CountDownLatch
}


func NewWorker(tasks chan Task, executor *Executor, cdl *sync.CountDownLatch) *Worker {
	return &Worker{
		tasks:    tasks,
		results:  make(chan Result),
		executor: executor,
		cdl:      cdl,
	}
}

func (w *Worker) AddTask(task Task) {
	w.tasks <- task
}

func (w *Worker) Run() {
	for {
		task, ok := <-w.tasks
		if !ok {
			if w.executor.Finished() {
				fmt.Println("Latch down by one!")
				w.cdl.CountDown()
				break
			} else {
				continue
			}
		}

		var execution Execution
		if task.GetType() {
			execution = NewReadExecution(task, w.executor)
		} else {
			execution = NewWriteExecution(task, w.executor)
		}

		result := execution.ExecuteTask()
		w.results <- result
	}
}

