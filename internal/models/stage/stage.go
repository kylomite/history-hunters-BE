package stage

import (
	"database/sql"
	"errors"
	"time"
)

type Stage struct {
	ID            int       `json:"id"`
	Title         string    `json:"title"`
	BackgroundImg string    `json:"background_img"`
	Difficulty    int       `json:"difficulty"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func NewStage(title, backgroundImg string, difficulty int) *Stage {
	return &Stage{
		Title:         title,
		BackgroundImg: backgroundImg,
		Difficulty:    difficulty,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
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

func GetAllStages(db *sql.DB) ([]Stage, error) {
	rows, err := db.Query(`SELECT id, title, background_img, difficulty, created_at, updated_at FROM stages`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var stages []Stage
	for rows.Next() {
		var s Stage
		if err := rows.Scan(&s.ID, &s.Title, &s.BackgroundImg, &s.Difficulty, &s.CreatedAt, &s.UpdatedAt); err != nil {
			return nil, err
		}
		stages = append(stages, s)
	}
	return stages, nil
}

func FindStageByID(db *sql.DB, id int) (*Stage, error) {
	var s Stage
	err := db.QueryRow(`
		SELECT id, title, background_img, difficulty, created_at, updated_at 
		FROM stages 
		WHERE id = $1`, id).
		Scan(&s.ID, &s.Title, &s.BackgroundImg, &s.Difficulty, &s.CreatedAt, &s.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("stage not found")
		}
		return nil, err
	}
	return &s, nil
}

func (s *Stage) Update(db *sql.DB) error {
	if s.Title == "" {
		return errors.New("title is required")
	}
	if s.BackgroundImg == "" {
		return errors.New("background image is required")
	}
	if s.Difficulty <= 0 {
		return errors.New("difficulty must be greater than 0")
	}

	query := `
		UPDATE stages
		SET title = $1, background_img = $2, difficulty = $3, updated_at = $4
		WHERE id = $5
		RETURNING updated_at;
	`
	err := db.QueryRow(query, s.Title, s.BackgroundImg, s.Difficulty, time.Now(), s.ID).Scan(&s.UpdatedAt)
	return err
}

func DeleteStage(db *sql.DB, id int) error {
	query := `DELETE FROM stages WHERE id = $1`
	_, err := db.Exec(query, id)
	return err
}