package main

import (
	"fmt"
	"sync"
)


type Executor struct {
	tasks     []Task
	workers   []*Worker
	lock      sync.Mutex
	isWriting bool
}

func (e *Executor) lock() {
	e.lock.Lock()
}

func (e *Executor) unlock() {
	e.lock.Unlock()
}

func (e *Executor) finished() bool {
	return len(e.tasks) == 0
}

func (e *Executor) setWriting(isWriting bool) {
	e.isWriting = isWriting
}

func (e *Executor) recover() Task {
	task := e.tasks[0]
	e.tasks = e.tasks[1:]
	return task
}

func NewExecutor(numberOfWorkers int, tasks []Task, cdl *sync.CountDownLatch) *Executor {
	lock := sync.Mutex{}
	executor := &Executor{
		tasks:     tasks,
		workers:   make([]*Worker, numberOfWorkers),
		lock:      lock,
		isWriting: false,
	}

	for i := 0; i < numberOfWorkers; i++ {
		w := &Worker{
			tasks:   make(chan Task),
			executor: executor,
			cdl:     cdl,
		}
		executor.workers[i] = w
		w.start()
	}

	return executor
}

func (e *Executor) dispatch() {
	for !e.finished() {
		for _, w := range e.workers {
			if e.finished() {
				break
			}
			task := e.recover()
			w.addTask(task)
		}
	}
}

func (e *Executor) getWriting() bool {
	return e.isWriting
}
