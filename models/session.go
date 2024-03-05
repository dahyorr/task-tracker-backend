package models

import (
	"time"

	"github.com/dahyorr/task-tracker-backend/database"
	"github.com/dahyorr/task-tracker-backend/utils"
	"github.com/gofiber/fiber/v2"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

type Session struct {
	Id        int64     `json:"id"`
	UserId    int64     `json:"user_id" db:"user_id"`
	Token     string    `json:"token"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	ExpiresAt time.Time `json:"expires_at" db:"expires_at"`
}

func (s *Session) Save() error {
	stmt := "INSERT INTO sessions (user_id, token, created_at, expires_at) VALUES ($1,$2,$3,$4) RETURNING id;"
	err := database.DB.QueryRow(stmt, s.UserId, s.Token, s.CreatedAt, s.ExpiresAt).Scan(&s.Id)
	if err != nil {
		return err
	}
	return nil
}

func (s *Session) IsExpired() bool {
	return s.ExpiresAt.Before(time.Now())
}

func NewSession(userId int64) (*Session, error) {
	token, err := gonanoid.New(20)
	if err != nil {
		return nil, err
	}
	session := Session{
		UserId:    userId,
		Token:     token,
		CreatedAt: time.Now(),
		ExpiresAt: time.Now().Add(utils.Config.SessionDuration),
	}
	err = session.Save()
	if err != nil {
		return nil, err
	}
	return &session, nil
}

func GetSessionByToken(token string) (*Session, error) {
	var session Session
	query := `SELECT * FROM sessions WHERE token=$1`
	err := database.DB.Get(&session, query, token)
	if err != nil {
		return nil, err
	}
	return &session, nil
}

func (s *Session) Delete() error {
	stmt := "DELETE FROM sessions WHERE id=$1"
	_, err := database.DB.Exec(stmt, s.Id)
	s.Token = ""
	if err != nil {
		return err
	}
	return nil
}

func (s *Session) Refresh() error {
	newToken, err := gonanoid.New(20)
	if err != nil {
		return err
	}
	s.ExpiresAt = time.Now().Add(utils.Config.SessionDuration)
	s.Token = newToken
	s.CreatedAt = time.Now()
	stmt := "UPDATE sessions SET expires_at=:expires_at, token=:token WHERE id=:id"
	_, err = database.DB.NamedExec(stmt, s)
	if err != nil {
		return err
	}
	return nil
}

func (s *Session) IsValid() bool {
	return !s.IsExpired()
}

func (s *Session) GetUser() (*User, error) {
	return GetUserById(s.UserId)
}

func (s *Session) ToCookie() *fiber.Cookie {
	cookie := new(fiber.Cookie)
	cookie.Name = utils.Config.SessionCookieName
	cookie.Value = s.Token
	cookie.Expires = s.ExpiresAt
	cookie.HTTPOnly = true
	cookie.Secure = true
	cookie.Domain = "localhost"
	cookie.SameSite = "Lax"
	return cookie
}
