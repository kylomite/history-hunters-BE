package stage

import (
	"database/sql"
	"errors"
	"time"
)

type Stage struct {
	ID           int
	Title        string
	BackgroundImg string
	Difficulty   int
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func NewStage(title, backgroundImg string, difficulty int) *Stage {
	return &Stage{
		Title:        title,
		BackgroundImg: backgroundImg,
		Difficulty:   difficulty,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
}

func (s *Stage) Save(db *sql.DB) error {
	if s.Title == "" {
		return errors.New("title is required")
	}
	if s.BackgroundImg == "" {
		return errors.New("background image is required")
	}
	if s.Difficulty <= 0 {
		return errors.New("difficulty is required")
	}

	query := `
		INSERT INTO stages (title, background_img, difficulty, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id;
	`
	err := db.QueryRow(query, s.Title, s.BackgroundImg, s.Difficulty, s.CreatedAt, s.UpdatedAt).Scan(&s.ID)
	if err != nil {
		return err
	}
	return nil
}

func (s *Stage) Delete(db *sql.DB) error {
    query := `DELETE FROM stages WHERE id = $1`
    _, err := db.Exec(query, s.ID)
    return err
}