package question

import (
	"database/sql"
	"errors"
	"time"
)

type Question struct {
	ID              int
	PlayerSessionID int
	QuestionText    string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

func NewQuestion(playerSessionID int, questionText string) *Question {
	return &Question{
		PlayerSessionID: playerSessionID,
		QuestionText:    questionText,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}
}

// Save persists the question to the database
func (q *Question) Save(db *sql.DB) error {
	if q.PlayerSessionID == 0 {
		return errors.New("player_session_id is required")
	}
	if q.QuestionText == "" {
		return errors.New("text is required")
	}

	query := `
		INSERT INTO questions (player_session_id, question_text, created_at, updated_at)
		VALUES ($1, $2, $3, $4)
		RETURNING id, created_at, updated_at;
	`

	err := db.QueryRow(query, q.PlayerSessionID, q.QuestionText, q.CreatedAt, q.UpdatedAt).
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