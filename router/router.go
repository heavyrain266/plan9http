package router

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"text/template"
)

var (
	buf   bytes.Buffer
	tasks = []map[string]string{
		{"id": "task-01", "title": "test", "details": "a test task", "status": "active"},
		{"id": "task-02", "title": "test 2", "details": "a test task 2", "status": "active"},
		{"id": "task-03", "title": "Test 3", "details": "a test task 3", "status": "active"},
		{"id": "task-04", "title": "Test 4", "details": "a test task 4", "status": "active"},
		{"id": "task-05", "title": "Test 5", "details": "a test task 5", "status": "active"},
		{"id": "task-06", "title": "Test 6", "details": "a test task 6", "status": "pending"},
		{"id": "task-07", "title": "Test 7", "details": "a test task 7", "status": "pending"},
		{"id": "task-08", "title": "Test 8", "details": "a test task 8", "status": "active"},
		{"id": "task-09", "title": "Test 9", "details": "a test task 9", "status": "inactive"},
		{"id": "task-10", "title": "Test 10", "details": "a test task 10", "status": "inactive"},
	}
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", taskListPage())

	return mux
}

func taskListPage() http.HandlerFunc {
	files := templLayout("./web/templates/index.gohtml")
	templ := template.Must(template.New("index").Funcs(defaultFuncs).ParseFiles(files...))

	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" || r.Method != http.MethodGet {
			notFoundPage(w, r)

			return
		}

		if err := templ.ExecuteTemplate(&buf, "base", map[string]interface{}{
			"Tasks": tasks,
		}); err != nil {
			fmt.Printf("ERR: %v\n", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)

			return
		}

		w.WriteHeader(http.StatusOK)
		io.Copy(w, &buf)
	}
}

var notFoundPage = func() http.HandlerFunc {
	files := templLayout("./web/templates/_notFound.gohtml")
	templ := template.Must(template.New("index").Funcs(defaultFuncs).ParseFiles(files...))

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var buf bytes.Buffer
		if err := templ.ExecuteTemplate(&buf, "base", nil); err != nil {
			fmt.Printf("ERR: %v\n", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		io.Copy(w, &buf)
	})
}()
