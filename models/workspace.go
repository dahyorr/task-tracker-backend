package models

import (
	"time"

	"github.com/dahyorr/task-tracker-backend/database"
)

type Workspace struct {
	Id        int64     `json:"id"`
	Name      string    `json:"name"`
	OwnerId   int64     `json:"owner_id" db:"owner_id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type WorkspaceFormData struct {
	Name string `json:"name" validate:"required"`
}

type WorkspaceMember struct {
	UserId      int `json:"user_id" db:"user_id"`
	WorkspaceId int `json:"workspace_id" db:"workspace_id"`
}

func (w *Workspace) Create() error {
	tx := database.DB.MustBegin()
	insertWorkspaceStmt := "INSERT INTO workspaces (name, owner_id) VALUES ($1,$2) RETURNING id, created_at;"
	insertWorkspaceMembersStmt := "INSERT INTO workspace_members (user_id, workspace_id) VALUES ($1,$2);"
	tx.QueryRow(insertWorkspaceStmt, w.Name, w.OwnerId).Scan(&w.Id, &w.CreatedAt)
	tx.Exec(insertWorkspaceMembersStmt, w.OwnerId, w.Id)
	return tx.Commit()
}

func (w *Workspace) Delete() error {
	// TODO: Add restrictions
	stmt := "DELETE FROM workspaces WHERE id = $1;"
	_, err := database.DB.Exec(stmt)
	return err
}

func (w *Workspace) AddUser(userId int64) error {
	stmt := "INSERT INTO workspace_members (user_id, workspace_id) VALUES ($1,$2);"
	_, err := database.DB.Exec(stmt, userId, w.Id)
	return err
}

func (w *Workspace) RemoveUser(userId int64) error {
	stmt := "DELETE FROM workspace_members WHERE user_id = $1 AND workspace_id = $2;"
	_, err := database.DB.Exec(stmt, userId, w.Id)
	return err
}

func GetWorkspacesByUserId(userId int64) ([]Workspace, error) {
	workspaces := []Workspace{}
	stmt := "SELECT w.id, w.name, w.owner_id, w.created_at FROM workspaces w JOIN workspace_members wm ON w.id = wm.workspace_id WHERE wm.user_id = $1;"
	err := database.DB.Select(&workspaces, stmt, userId)
	if err != nil {
		return nil, err
	}
	return workspaces, nil
}

func GetWorkspaceById(id int64) (Workspace, error) {
	workspace := Workspace{}
	stmt := "SELECT id, name, owner_id, created_at FROM workspaces WHERE id = $1 ;"
	err := database.DB.Get(&workspace, stmt, id)
	return workspace, err
}

func GetWorkspaceMembers(workspaceId int64) ([]WorkspaceMember, error) {
	members := []WorkspaceMember{}
	stmt := "SELECT user_id, workspace_id FROM workspace_members WHERE workspace_id = $1;"
	err := database.DB.Select(&members, stmt, workspaceId)
	return members, err
}
