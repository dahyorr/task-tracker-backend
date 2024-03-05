package models

import (
	"time"

	"github.com/dahyorr/task-tracker-backend/database"
)

type Task struct {
	Id          int64     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedBy   int64     `json:"created_by" db:"created_by"`
	WorkspaceId int64     `json:"workspace_id" db:"workspace_id"`
	DueDate     time.Time `json:"due_date" db:"due_date"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

type TaskForm struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	WorkspaceId int64     `json:"workspace_id" db:"workspace_id"`
	DueDate     time.Time `json:"due_date" db:"due_date"`
}

func (t *Task) Create() error {
	err := database.DB.QueryRow("INSERT INTO tasks (name, description, status, created_by, workspace_id, due_date) VALUES ($1,$2,$3,$4,$5,$6) RETURNING id, created_at, updated_at", t.Name, t.Description, t.Status, t.CreatedBy, t.WorkspaceId, t.DueDate).Scan(&t.Id, &t.CreatedAt, &t.UpdatedAt)
	return err
}

func (t *Task) Update() error {
	_, err := database.DB.Exec("UPDATE tasks SET name=$1, description=$2, status=$3, due_date=$4 WHERE id=$5", t.Name, t.Description, t.Status, t.DueDate, t.Id)
	return err
}

func (t *Task) UpdateStatus(status string) error {
	_, err := database.DB.Exec("UPDATE tasks SET status=$1 WHERE id=$2", status, t.Id)
	if err != nil {
		return err
	}
	t.Status = status
	return nil
}

func (t *Task) Delete() error {
	_, err := database.DB.Exec("DELETE FROM tasks WHERE id = $1", t.Id)
	return err
}

func GetTasksByWorkspaceId(workspaceId int64) ([]Task, error) {
	tasks := []Task{}
	err := database.DB.Select(&tasks, "SELECT * FROM tasks WHERE workspace_id = $1", workspaceId)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func GetTaskById(taskId int64) (*Task, error) {
	task := &Task{}
	err := database.DB.Get(task, "SELECT * FROM tasks WHERE id = $1", taskId)
	if err != nil {
		return nil, err
	}
	return task, nil
}
