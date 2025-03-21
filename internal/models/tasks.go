package models

import "time"

type Task struct {
	Id      int       `json:"id"`
	Title   string    `json:"title"`
	Content string    `json:"content"`
	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
}

type TaskStore interface {
	GetAll() ([]Task, error)
	Get(id int) (Task, error)
	Create(title, content string) (Task, error)
	Update(id int, title, content string) error
	Delete(id int) error
}

type TaskStoreInMemory struct {
	tasks map[int]Task
}

func NewTaskStoreInMemory() *TaskStoreInMemory {
	return &TaskStoreInMemory{
		tasks: map[int]Task{},
	}
}

func (s *TaskStoreInMemory) GetAll() ([]Task, error) {
	tasks := make([]Task, 0, len(s.tasks))

	for _, t := range s.tasks {
		tasks = append(tasks, t)
	}

	return tasks, nil
}

func (s *TaskStoreInMemory) Get(id int) (Task, error) {
	t, ok := s.tasks[id]
	if !ok {
		return Task{}, ErrRecordNotFound
	}

	return t, nil
}

func (s *TaskStoreInMemory) Create(title, content string) (Task, error) {
	t := Task{
		Id:      len(s.tasks) + 1,
		Title:   title,
		Content: content,
		Created: time.Now(),
		Updated: time.Now(),
	}

	s.tasks[t.Id] = t
	return t, nil
}

func (s *TaskStoreInMemory) Update(id int, title, content string) error {
	t, ok := s.tasks[id]
	if !ok {
		return ErrRecordNotFound
	}

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
	if _, ok := s.tasks[id]; !ok {
		return ErrRecordNotFound
	}

	delete(s.tasks, id)
	return nil
}
