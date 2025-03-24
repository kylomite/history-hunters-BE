package models

import (
	"database/sql"
	"errors"
	"time"
)

// Player represents the player model.
type Player struct {
	ID             int
	Email          string
	PasswordDigest string
	Avatar         string
	Score          int
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func NewPlayer(email, passwordDigest, avatar string) *Player {
	return &Player{
		Email:          email,
		PasswordDigest: passwordDigest,
		Avatar:         avatar,
		Score:          0,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}
}

func (p *Player) Save(db *sql.DB) error {
	query := `
		INSERT INTO players (email, password_digest, avatar, score, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id;
	`
	err := db.QueryRow(query, p.Email, p.PasswordDigest, p.Avatar, p.Score, p.CreatedAt, p.UpdatedAt).Scan(&p.ID)
	if err != nil {
		if err.Error() == `pq: duplicate key value violates unique constraint "players_email_key"` {
			return errors.New("email already exists")
		}
		return err
	}
	return nil
}

func EmailExists(db *sql.DB, email string) (bool, error) {
	query := `SELECT COUNT(*) FROM players WHERE email = $1`
	var count int
	err := db.QueryRow(query, email).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}