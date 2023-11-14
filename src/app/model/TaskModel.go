package model

import (
	"time"
)

type Task struct {
	TaskNumber       int       `json:"tasknumber"`
	Heading          string    `json:"heading"`
	Data             string    `json:"data,omitempty"`
	Creationdatetime time.Time `json:"creationdatetime"`
	Updationdatetime time.Time `json:"updationdatetime,omitempty"`
}

var taskNumberCounter int = 1

func (t *Task) New() {
	t.TaskNumber = taskNumberCounter
	t.Creationdatetime = time.Now()
	t.Updationdatetime = time.Now()
	taskNumberCounter += 1
}

func (t *Task) UpdateTask(requestTask Task) {
	if requestTask.Data != "" {
		t.Data = requestTask.Data
	}
	if requestTask.Heading != "" {
		t.Heading = requestTask.Heading
	}
	t.Updationdatetime = time.Now()
}
