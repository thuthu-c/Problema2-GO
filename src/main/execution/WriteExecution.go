package execution

import (
	"fmt"
	"time"
)

type WriteExecution struct {
	task     taskexecutor.Task
	result   taskexecutor.Result
	executor *taskexecutor.Executor
	begin    time.Time
}

func NewWriteExecution(t taskexecutor.Task, e *taskexecutor.Executor) *WriteExecution {
	return &WriteExecution{
		begin:    time.Now(),
		task:     t,
		executor: e,
	}
}

func (e *WriteExecution) Hold() {
	time.Sleep(time.Duration(e.task.GetCost()) * time.Second)
}

func (e *WriteExecution) Read() int {
	return taskexecutor.GetFileValue()
}

func (e *WriteExecution) Fill(value int) {
	e.result = taskexecutor.NewResult(e.task.GetID(), value, time.Since(e.begin))
}


func (e *WriteExecution) sum() int{

	return task.getValue() + read() 
}

func (e *WritwExecution) write(int value){
	FileHandler.writeToFile(value)
}

func (e *WriteExecution) ExecuteTask(Task task) taskexecutor.Result {
	executor.setWriting(true)
	executor.lock()

	hold()

	var value int 
	value := sum()

	write(value)
	fill(value)

	executor.unlock()
	executor.setWriting(false)

	return result
}
