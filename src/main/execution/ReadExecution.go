package execution

import (
	"fmt"
	"time"
)

type ReadExecution struct {
	task     taskexecutor.Task
	result   taskexecutor.Result
	executor *taskexecutor.Executor
	begin    time.Time
}

func NewReadExecution(t taskexecutor.Task, e *taskexecutor.Executor) *ReadExecution {
	return &ReadExecution{
		begin:    time.Now(),
		task:     t,
		executor: e,
	}
}

func (e *ReadExecution) Hold() {
	time.Sleep(time.Duration(e.task.GetCost()) * time.Second)
}

func (e *ReadExecution) Read() int {
	return taskexecutor.GetFileValue()
}

func (e *ReadExecution) Fill(value int) {
	e.result = taskexecutor.NewResult(e.task.GetID(), value, time.Since(e.begin))
}


func (e *ReadExecution) ExecuteTask(Task task) taskexecutor.Result {
	for executor.getWriting(){}
	hold()
	fill(read())

	return result
}
