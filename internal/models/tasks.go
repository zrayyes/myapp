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

type TaskStoreInMemory struct {
	tasks []Task
}

func NewTaskStoreInMemory() *TaskStoreInMemory {
	return &TaskStoreInMemory{
		tasks: []Task{},
	}
}

func (s *TaskStoreInMemory) GetAll() ([]Task, error) {
	return s.tasks, nil
}

func (s *TaskStoreInMemory) Get(id int) (Task, error) {
	return s.tasks[id], nil
}

func (s *TaskStoreInMemory) Create(t Task) (int, error) {
	s.tasks = append(s.tasks, t)
	return len(s.tasks) - 1, nil
}

func (s *TaskStoreInMemory) Update(id int, t Task) error {
	s.tasks[id] = t
	return nil
}

func (s *TaskStoreInMemory) Delete(id int) error {
	s.tasks = append(s.tasks[:id], s.tasks[id+1:]...)
	return nil
}
