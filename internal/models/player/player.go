package player

import (
	"database/sql"
	"errors"
	"time"
)

type Player struct {
	ID             int       `json:"id"`
	Email          string    `json:"email"`
	PasswordDigest string    `json:"password_digest"`
	Avatar         string    `json:"avatar"`
	Score          int       `json:"score"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func (p *Player) Save(db *sql.DB) error {
	query := `
		INSERT INTO players (email, password_digest, avatar, score, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id;
	`
	err := db.QueryRow(query, p.Email, p.PasswordDigest, p.Avatar, p.Score, time.Now(), time.Now()).Scan(&p.ID)
	if err != nil {
		return err
	}
	return nil
}

func GetAllPlayers(db *sql.DB) ([]Player, error) {
	rows, err := db.Query(`SELECT id, email, password_digest, avatar, score, created_at, updated_at FROM players`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var players []Player
	for rows.Next() {
		var p Player
		if err := rows.Scan(&p.ID, &p.Email, &p.PasswordDigest, &p.Avatar, &p.Score, &p.CreatedAt, &p.UpdatedAt); err != nil {
			return nil, err
		}
		players = append(players, p)
	}
	return players, nil
}

func FindPlayerByID(db *sql.DB, id int) (*Player, error) {
	var p Player
	err := db.QueryRow(`SELECT id, email, password_digest, avatar, score, created_at, updated_at FROM players WHERE id = $1`, id).
		Scan(&p.ID, &p.Email, &p.PasswordDigest, &p.Avatar, &p.Score, &p.CreatedAt, &p.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("player not found")
		}
		return nil, err
	}
	return &p, nil
}

func (p *Player) Update(db *sql.DB) error {
	_, err := db.Exec(`
		UPDATE players
		SET email = $1, password_digest = $2, avatar = $3, score = $4, updated_at = $5
		WHERE id = $6`,
		p.Email, p.PasswordDigest, p.Avatar, p.Score, time.Now(), p.ID)

	return err
}

func DeletePlayer(db *sql.DB, id int) error {
	_, err := db.Exec(`DELETE FROM players WHERE id = $1`, id)
	return err
}