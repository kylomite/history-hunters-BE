package answer

import (
	"database/sql"
	"errors"
	"time"
)

type Answer struct {
	ID         int
	QuestionID int
	AnswerText string
	Correct    bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func NewAnswer(questionID int, answerText string, correct bool) *Answer {
	return &Answer{
		QuestionID: questionID,
		AnswerText: answerText,
		Correct:    correct,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
}

func (a *Answer) Save(db *sql.DB) error {
	if a.QuestionID == 0 {
		return errors.New("question_id is required")
	}
	if a.AnswerText == "" {
		return errors.New("answer_text is required")
	}

	query := `
		INSERT INTO answers (question_id, answer_text, correct, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, created_at, updated_at;
	`

	err := db.QueryRow(query, a.QuestionID, a.AnswerText, a.Correct, a.CreatedAt, a.UpdatedAt).
		Scan(&a.ID, &a.CreatedAt, &a.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}

func FindByID(db *sql.DB, answerID int) (*Answer, error) {
	query := `
		SELECT id, question_id, answer_text, correct, created_at, updated_at
		FROM answers
		WHERE id = $1
	`

	row := db.QueryRow(query, answerID)

	a := &Answer{}
	err := row.Scan(&a.ID, &a.QuestionID, &a.AnswerText, &a.Correct, &a.CreatedAt, &a.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("answer not found")
		}
		return nil, err
	}

	return a, nil
}

func (a *Answer) Delete(db *sql.DB) error {
	query := `DELETE FROM answers WHERE id = $1`
	_, err := db.Exec(query, a.ID)
	return err
}