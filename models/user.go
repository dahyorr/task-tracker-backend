package models

import (
	"database/sql"
	"fmt"

	"dayo.dev/task-tracker/database"
	"dayo.dev/task-tracker/shared"
	"dayo.dev/task-tracker/utils"
)

type User struct {
	Id             int64          `json:"id"`
	FirstName      string         `json:"first_name" db:"first_name"`
	LastName       string         `json:"last_name" db:"last_name"`
	Email          string         `json:"email" validate:"required,email"`
	Username       sql.NullString `json:"-" `
	HashedPassword string         `db:"password" json:"-"`
}

type UserFormData struct {
	FirstName string `json:"first_name" db:"first_name" validate:"required"`
	LastName  string `json:"last_name" db:"last_name" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `validate:"required"`
}

type UserLoginFormData struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required"`
}

func (u *UserFormData) CreateUser() (*User, error) {
	user := User{
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
	}
	err := user.SetPassword(u.Password)
	if err != nil {
		return nil, err
	}
	err = user.Save()
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *User) Save() error {
	stmt := "INSERT INTO users (first_name, last_name, email, password) VALUES ($1,$2,$3,$4) RETURNING id;"
	return database.DB.QueryRow(stmt, u.FirstName, u.LastName, u.Email, u.HashedPassword).Scan(&u.Id)
}

func (u *User) SetPassword(password string) error {
	hashedPassword, err := utils.HashPassword(password)
	u.HashedPassword = hashedPassword
	return err
}

func (u *User) Update() error {
	stmt := `UPDATE users SET first_name=:first_name, last_name=:last_name, email=:email, password=:password WHERE id=:id`
	_, err := database.DB.NamedExec(stmt, u)
	if err != nil {
		return err
	}
	return nil
}

func (u *User) UpdatePassword(password string) error {
	err := u.SetPassword(password)
	if err != nil {
		return err
	}
	return u.Update()
}

func GetUserByEmail(email string) (*User, error) {
	var user User
	query := `SELECT * FROM users WHERE email=$1;`
	fmt.Println()
	err := database.DB.Get(&user, query, email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUserById(id int64) (*User, error) {
	var user User
	query := `SELECT * FROM users WHERE id=$1;`
	err := database.DB.Get(&user, query, id)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *User) ValidatePassword(password string) error {
	invalidError := shared.ErrInvalidEmailOrPassword
	passwordIsValid := utils.ComparePassword(password, u.HashedPassword)
	if !passwordIsValid {
		return invalidError
	}
	return nil
}
