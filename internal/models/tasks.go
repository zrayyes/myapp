package models

import "time"

type Task struct {
	Title   string    `json:"title"`
	Content string    `json:"content"`
	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
}

type TaskStore interface {
	GetAll() ([]Task, error)
	Get(id int) (Task, error)
	Create(title, content string) (int, error)
	Update(id int, title, content string) error
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

func (s *TaskStoreInMemory) Create(title, content string) (int, error) {
	t := Task{
		Title:   title,
		Content: content,
		Created: time.Now(),
		Updated: time.Now(),
	}

	s.tasks = append(s.tasks, t)
	return len(s.tasks) - 1, nil
}

func (s *TaskStoreInMemory) Update(id int, title, content string) error {
	t := s.tasks[id]

	if title != "" {
		t.Title = title
	}

	if content != "" {
		t.Content = content
	}

	t.Updated = time.Now()

	s.tasks[id] = t
	return nil
}

func (s *TaskStoreInMemory) Delete(id int) error {
	s.tasks = append(s.tasks[:id], s.tasks[id+1:]...)
	return nil
}
