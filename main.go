package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

type Task struct {
	ID string `json:"id"` // uuid

	Name string `json:"name"`

	StartTime time.Time `json:"startTime"` // 2023-01-05T01:00:00Z

	// The end time of the shift
	EndTime time.Time `json:"endTime"` // 2023-01-05T09:00:00Z
}

type CreateTaskRequest struct {
	StartTime *string `json:"startTime" validate:"required"`
	EndTime   *string `json:"endTime" validate:"required"`
	Name      *string `json:"name" validate:"required"`
}

func main() {
	// fmt.Println(task.StartTime)
	http.HandleFunc("/create", create_task)
	http.HandleFunc("/tasks", get_tasks)
	http.ListenAndServe(":8090", nil)

}

func print_tasks() {
	for _, task := range tasks {
		fmt.Println(task)
	}
}

var tasks []Task
var last_id int = 0

func create_task(w http.ResponseWriter, r *http.Request) {
	var req CreateTaskRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		fmt.Println("decode error")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	parsed_start_time, err := time.Parse("2006-01-02T15:04:05Z07:00", *req.StartTime)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	parsed_end_time, err := time.Parse("2006-01-02T15:04:05Z07:00", *req.EndTime)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var new_task Task = Task{
		ID:        strconv.Itoa(last_id + 1),
		Name:      *req.Name,
		StartTime: parsed_start_time,
		EndTime:   parsed_end_time,
	}
	last_id = last_id + 1
	w.Header().Set("Content-Type", "application/json")

	// task := &Task{ID: id, StartTime: time.Now(), EndTime: time.Now()}

	tasks = append(tasks, new_task)
	print_tasks()
	json.NewEncoder(w).Encode(new_task)
}
func get_tasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(tasks)
}
