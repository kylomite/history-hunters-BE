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

func FindByID(db *sql.DB, questionID, playerID int) (*Question, error) {
	query := `
		SELECT q.id, q.player_session_id, q.question_text, q.created_at, q.updated_at
		FROM questions q
		INNER JOIN player_sessions ps ON q.player_session_id = ps.id
		WHERE q.id = $1 AND ps.player_id = $2
	`

	row := db.QueryRow(query, questionID, playerID)

	q := &Question{}
	err := row.Scan(&q.ID, &q.PlayerSessionID, &q.QuestionText, &q.CreatedAt, &q.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("question not found or unauthorized access")
		}
		return nil, err
	}

	return q, nil
}

func (p *Question) Delete(db *sql.DB) error {
    query := `DELETE FROM questions WHERE id = $1`
    _, err := db.Exec(query, p.ID)
    return err
}