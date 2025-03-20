package models

import "time"

type Task struct {
	Title   string
	Content string
	Created time.Time
	Updated time.Time
}

type TaskStore interface {
	GetAll() ([]Task, error)
	Get(id int) (Task, error)
	Create(t Task) (int, error)
	Update(id int, t Task) error
	Delete(id int) error
}
