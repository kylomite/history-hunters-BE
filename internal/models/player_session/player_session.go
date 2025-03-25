package player_session

import (
	"database/sql"
	"errors"
	"time"
)

type PlayerSession struct {
	ID        int       `json:"id"`
	PlayerID  int       `json:"player_id"`
	StageID   int       `json:"stage_id"`
	Lives     int       `json:"lives"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
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
	return err
}

func GetAllByPlayerID(db *sql.DB, playerID int) ([]PlayerSession, error) {
	query := `
		SELECT id, player_id, stage_id, lives, created_at, updated_at
		FROM player_sessions
		WHERE player_id = $1
	`

	rows, err := db.Query(query, playerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sessions []PlayerSession
	for rows.Next() {
		var ps PlayerSession
		if err := rows.Scan(&ps.ID, &ps.PlayerID, &ps.StageID, &ps.Lives, &ps.CreatedAt, &ps.UpdatedAt); err != nil {
			return nil, err
		}
		sessions = append(sessions, ps)
	}

	return sessions, nil
}

func GetByID(db *sql.DB, sessionID int) (*PlayerSession, error) {
	query := `
		SELECT id, player_id, stage_id, lives, created_at, updated_at
		FROM player_sessions
		WHERE id = $1
	`

	var ps PlayerSession
	err := db.QueryRow(query, sessionID).Scan(
		&ps.ID, &ps.PlayerID, &ps.StageID, &ps.Lives, &ps.CreatedAt, &ps.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("player session not found")
		}
		return nil, err
	}

	return &ps, nil
}

func (ps *PlayerSession) Update(db *sql.DB) error {
	if ps.Lives <= 0 {
		return errors.New("lives must be greater than 0")
	}

	query := `
		UPDATE player_sessions
		SET lives = $1, updated_at = $2
		WHERE id = $3
		RETURNING updated_at;
	`

	err := db.QueryRow(query, ps.Lives, time.Now(), ps.ID).Scan(&ps.UpdatedAt)
	return err
}

func (ps *PlayerSession) Delete(db *sql.DB) error {
	query := `DELETE FROM player_sessions WHERE id = $1`
	_, err := db.Exec(query, ps.ID)
	return err
}