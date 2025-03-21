package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/zrayyes/myapp/internal/models"
)

func UnmarshalResponse(t *testing.T, r *httptest.ResponseRecorder, want interface{}) {
	t.Helper()

	err := json.Unmarshal(r.Body.Bytes(), want)
	if err != nil {
		t.Fatalf("could not unmarshal response: %v", err)
	}
}

func TestGetTasks(t *testing.T) {
	app := newApplication(models.NewTaskStoreInMemory())

	app.taskStore.Create("Salt", "This is Salt's task")

	t.Run("returns Pepper's score", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/tasks", nil)
		response := httptest.NewRecorder()

		app.getTasks(response, request)

		want := &models.Task{
			Id:      1,
			Title:   "Salt",
			Content: "This is Salt's task",
		}

		rTasks := []models.Task{}

		UnmarshalResponse(t, response, &rTasks)

		if len(rTasks) != 1 {
			t.Fatalf("got %d, want %d", len(rTasks), 1)
		}

		if rTasks[0].Id != want.Id {
			t.Errorf("got %d, want %d", rTasks[0].Id, want.Id)
		}

		if rTasks[0].Title != want.Title {
			t.Errorf("got %s, want %s", rTasks[0].Title, want.Title)
		}

		if rTasks[0].Content != want.Content {
			t.Errorf("got %s, want %s", rTasks[0].Content, want.Content)
		}

	})
}
