package model

import (
	"time"
)

type Task struct {
	TaskNumber       int       `json:"tasknumber"`
	Heading          string    `json:"heading"`
	Data             string    `json:"data"`
	Creationdatetime time.Time `json:"creationdatetime"`
}

func (t *Task) New() {
	t.Creationdatetime = time.Now()
}
