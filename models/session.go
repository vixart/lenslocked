package models

import (
	"database/sql"
	"fmt"

	"github.com/sixsat/lenslocked/rand"
)

type Session struct {
	ID     int
	UserID int
	// Token is only set when creating a new session.
	// It will be left empty when look up a session,
	// as we only store the hash of a session token in the database.
	Token     string
	TokenHash string
}

type SessionService struct {
	DB *sql.DB
}

func (ss *SessionService) Create(userID int) (*Session, error) {
	token, err := rand.SessionToken()
	if err != nil {
		return nil, fmt.Errorf("create: %w", err)
	}
	// TODO: Hash the session token
	session := Session{
		UserID: userID,
		Token:  token,
		// TODO: Set the TokenHash
	}
	// TODO: Store session in the DB
	return &session, nil
}

func (ss *SessionService) User(token string) (*User, error) {
	// TODO: Implement SessionService.User
	return nil, nil
}
