package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/zrayyes/myapp/internal/models"
)

func UnmarshalResponse(t *testing.T, r *http.Response, want interface{}) {
	t.Helper()

	if err := json.NewDecoder(r.Body).Decode(want); err != nil {
		t.Fatalf("could not unmarshal JSON response: %v", err)
	}
}

func setupTestApp() *application {
	app := newApplication(models.NewTaskStoreInMemory())
	return app
}

func TestGetTasks(t *testing.T) {
	app := setupTestApp()

	ts := httptest.NewServer(app.routes())
	defer ts.Close()

	app.taskStore.Create("Salt", "This is Salt's task")
	app.taskStore.Create("Pepper", "This is Pepper's task")

	tests := []struct {
		name           string
		expectedTasks  []models.Task
		expectedStatus int
	}{
		{
			name: "GET all tasks",
			expectedTasks: []models.Task{
				{Id: 1, Title: "Salt", Content: "This is Salt's task"},
				{Id: 2, Title: "Pepper", Content: "This is Pepper's task"},
			},
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			response, err := http.Get(ts.URL + "/tasks")
			if err != nil {
				t.Fatalf("could not send GET request: %v", err)
			}

			if response.StatusCode != tt.expectedStatus {
				t.Fatalf("got status %d, want %d", response.StatusCode, tt.expectedStatus)
			}

			var gotTasks []models.Task
			UnmarshalResponse(t, response, &gotTasks)

			if len(gotTasks) != len(tt.expectedTasks) {
				t.Fatalf("got %d tasks, want %d", len(gotTasks), len(tt.expectedTasks))
			}

			for i, got := range gotTasks {
				want := tt.expectedTasks[i]
				if got.Id != want.Id || got.Title != want.Title || got.Content != want.Content {
					t.Errorf("got %+v, want %+v", got, want)
				}
			}
		})
	}
}

func TestGetTask(t *testing.T) {
	app := setupTestApp()

	ts := httptest.NewServer(app.routes())
	defer ts.Close()

	app.taskStore.Create("Salt", "This is Salt's task")

	tests := []struct {
		name           string
		taskID         string
		expectedTask   *models.Task
		expectedStatus int
	}{
		{
			name:           "GET existing task",
			taskID:         "1",
			expectedTask:   &models.Task{Id: 1, Title: "Salt", Content: "This is Salt's task"},
			expectedStatus: http.StatusOK,
		},
		{
			name:           "GET non-existing task",
			taskID:         "999",
			expectedTask:   nil,
			expectedStatus: http.StatusNotFound,
		},
		{
			name:           "GET invalid task ID",
			taskID:         "abc",
			expectedTask:   nil,
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			response, err := http.Get(ts.URL + "/tasks/" + tt.taskID)
			if err != nil {
				t.Fatalf("could not send GET request: %v", err)
			}

			if response.StatusCode != tt.expectedStatus {
				t.Fatalf("got status %d, want %d", response.StatusCode, tt.expectedStatus)
			}

			if tt.expectedTask != nil {
				var gotTask models.Task
				UnmarshalResponse(t, response, &gotTask)

				if gotTask.Id != tt.expectedTask.Id || gotTask.Title != tt.expectedTask.Title || gotTask.Content != tt.expectedTask.Content {
					t.Errorf("got %+v, want %+v", gotTask, tt.expectedTask)
				}
			}
		})
	}
}

func TestCreateTask(t *testing.T) {
	app := setupTestApp()

	ts := httptest.NewServer(app.routes())
	defer ts.Close()

	tests := []struct {
		name           string
		task           models.Task
		expectedTask   *models.Task
		expectedStatus int
	}{
		{
			name: "POST new task",
			task: models.Task{Title: "Salt", Content: "This is Salt's task"},
			expectedTask: &models.Task{
				Id:      1,
				Title:   "Salt",
				Content: "This is Salt's task",
			},
			expectedStatus: http.StatusCreated,
		},
		{
			name:           "POST empty task",
			task:           models.Task{},
			expectedTask:   nil,
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body, err := json.Marshal(tt.task)
			if err != nil {
				t.Fatalf("could not marshal task: %v", err)
			}

			response, err := http.Post(ts.URL+"/tasks", "application/json", bytes.NewReader(body))
			if err != nil {
				t.Fatalf("could not send POST request: %v", err)
			}

			if response.StatusCode != tt.expectedStatus {
				t.Fatalf("got status %d, want %d", response.StatusCode, tt.expectedStatus)
			}

			if tt.expectedTask != nil {
				var gotTask models.Task
				UnmarshalResponse(t, response, &gotTask)

				if gotTask.Id != tt.expectedTask.Id || gotTask.Title != tt.expectedTask.Title || gotTask.Content != tt.expectedTask.Content {
					t.Errorf("got %+v, want %+v", gotTask, tt.expectedTask)
				}
			}
		})
	}
}
