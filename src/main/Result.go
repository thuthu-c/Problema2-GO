package main

import (
	"time"
)

type Result struct{
	id      int
    result  int
    time    time.Duration
}

func NewResult(id int, result int, time time.Duration) *Result {
	return &Result{
		id:    		id,
		result:  	result,
		time: 		time,
	}
}


