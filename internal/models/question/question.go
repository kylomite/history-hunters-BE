package question

import (
	"database/sql"
	"errors"
	"time"
)

type Question struct {
	ID              int
	PlayerSessionID int
	Text            string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

func NewQuestion(playerSessionID int, text string) *Question {
	return &Question{
		PlayerSessionID: playerSessionID,
		Text:            text,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}
}

// Save persists the question to the database
func (q *Question) Save(db *sql.DB) error {
	if q.PlayerSessionID == 0 {
		return errors.New("player_session_id is required")
	}
	if q.Text == "" {
		return errors.New("text is required")
	}

	query := `
		INSERT INTO questions (player_session_id, text, created_at, updated_at)
		VALUES ($1, $2, $3, $4)
		RETURNING id, created_at, updated_at;
	`

	err := db.QueryRow(query, q.PlayerSessionID, q.Text, q.CreatedAt, q.UpdatedAt).
		Scan(&q.ID, &q.CreatedAt, &q.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (p *Question) Delete(db *sql.DB) error {
    query := `DELETE FROM questions WHERE id = $1`
    _, err := db.Exec(query, p.ID)
    return err
}