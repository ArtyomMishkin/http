package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"example.com/http/internal/storage"
)

func TestCreateTask(t *testing.T) {
	store := storage.NewMemoryStore()
	h := NewHandlers(store)

	payload := `{"title":"Test task"}`
	req := httptest.NewRequest(http.MethodPost, "/tasks", strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	h.CreateTask(w, req)

	res := w.Result()
	if res.StatusCode != http.StatusCreated {
		t.Fatalf("expected status 201, got %d", res.StatusCode)
	}

	var tsk storage.Task
	if err := json.NewDecoder(res.Body).Decode(&tsk); err != nil {
		t.Fatal("failed to decode response:", err)
	}
	if tsk.Title != "Test task" || tsk.Done {
		t.Fatalf("unexpected task data: %+v", tsk)
	}
}

func TestCreateTask_EmptyTitle(t *testing.T) {
	store := storage.NewMemoryStore()
	h := NewHandlers(store)

	payload := `{"title":""}`
	req := httptest.NewRequest(http.MethodPost, "/tasks", strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	h.CreateTask(w, req)

	res := w.Result()
	if res.StatusCode != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d", res.StatusCode)
	}

	var errResp ErrorResponse
	if err := json.NewDecoder(res.Body).Decode(&errResp); err != nil {
		t.Fatal("failed to decode response:", err)
	}
	if errResp.Error != "title is required" {
		t.Fatalf("unexpected error message: %s", errResp.Error)
	}
}

func TestListTasks(t *testing.T) {
	store := storage.NewMemoryStore()
	h := NewHandlers(store)

	// Pre-populate with test data
	store.Create("First task")
	store.Create("Second task")

	req := httptest.NewRequest(http.MethodGet, "/tasks", nil)
	w := httptest.NewRecorder()

	h.ListTasks(w, req)

	res := w.Result()
	if res.StatusCode != http.StatusOK {
		t.Fatalf("expected status 200, got %d", res.StatusCode)
	}

	var tasks []storage.Task
	if err := json.NewDecoder(res.Body).Decode(&tasks); err != nil {
		t.Fatal("failed to decode response:", err)
	}
	if len(tasks) != 2 {
		t.Fatalf("expected 2 tasks, got %d", len(tasks))
	}
}

func TestListTasks_WithFilter(t *testing.T) {
	store := storage.NewMemoryStore()
	h := NewHandlers(store)

	// Добавляем несколько тестов
	store.Create("Buy milk")
	store.Create("Write code")
	store.Create("Drink Beer")

	req := httptest.NewRequest(http.MethodGet, "/tasks?q=code", nil)
	w := httptest.NewRecorder()

	h.ListTasks(w, req)

	res := w.Result()
	if res.StatusCode != http.StatusOK {
		t.Fatalf("expected status 200, got %d", res.StatusCode)
	}

	var tasks []storage.Task
	if err := json.NewDecoder(res.Body).Decode(&tasks); err != nil {
		t.Fatal("failed to decode response:", err)
	}
	if len(tasks) != 1 {
		t.Fatalf("expected 1 task, got %d", len(tasks))
	}
	if tasks[0].Title != "Write code" {
		t.Fatalf("unexpected task title: %s", tasks[0].Title)
	}
}
func TestCreateTask_InvalidLength(t *testing.T) {
	store := storage.NewMemoryStore()
	h := NewHandlers(store)

	// Тест слишком короткого заголовка
	payload := `{"title":"ab"}`
	req := httptest.NewRequest(http.MethodPost, "/tasks", strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	h.CreateTask(w, req)
	res := w.Result()
	if res.StatusCode != http.StatusUnprocessableEntity {
		t.Fatalf("expected status 422, got %d", res.StatusCode)
	}

	// Тест слишком длинного заголовка
	longTitle := strings.Repeat("a", 141)
	payload = `{"title":"` + longTitle + `"}`
	req = httptest.NewRequest(http.MethodPost, "/tasks", strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()

	h.CreateTask(w, req)

	res = w.Result()
	if res.StatusCode != http.StatusUnprocessableEntity {
		t.Fatalf("expected status 422, got %d", res.StatusCode)
	}
}
