package main

import "sync"


type Task struct {
	id              	int
	cost    			float64
	types     			bool
	value			    int8
}

func NewTask(id int, cost float64, types bool, value int8) *Task {
	return &Task{
		id:    		id,
		cost:  		cost,
		types: 		types,
		value:      value,
	}
}


