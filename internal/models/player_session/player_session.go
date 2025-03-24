package models

import (
	"database/sql"
	"errors"
	"time"
)

type PlayerSession struct {
	ID        int
	PlayerID  int
	StageID   int
	Lives     int
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewPlayerSession(playerID, stageID, lives int) *PlayerSession {
	return &PlayerSession{
		PlayerID:  playerID,
		StageID:   stageID,
		Lives:     lives,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (ps *PlayerSession) Save(db *sql.DB) error {
	if ps.PlayerID == 0 {
		return errors.New("player_id is required")
	}
	if ps.StageID == 0 {
		return errors.New("stage_id is required")
	}
	if ps.Lives <= 0 {
		return errors.New("lives must be greater than 0")
	}

	query := `
		INSERT INTO player_sessions (player_id, stage_id, lives, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id;
	`
	err := db.QueryRow(query, ps.PlayerID, ps.StageID, ps.Lives, ps.CreatedAt, ps.UpdatedAt).Scan(&ps.ID)
	if err != nil {
		return err
	}
	return nil
}

func (ps *PlayerSession) Delete(db *sql.DB) error {
    query := `DELETE FROM player_sessions WHERE id = $1`
    _, err := db.Exec(query, ps.ID)
    return err
}