package store

import (
	"database/sql"
	"errors"
	_ "github.com/lib/pq"
	"time"
)

type Competence struct {
	Id         int       `json:"id"`
	Nome       string    `json:"nome"`
	DtCriacao  time.Time `json:"dt_criacao"`
	Processada bool      `json:"processada"`
}

type CompetenceStore interface {
	GetMaxCompetence() (string, error)
}

type SQLCompetenceStore struct {
	DB *sql.DB
}

func (s *SQLCompetenceStore) GetMaxCompetence() (string, error) {
	var competence Competence

	row := s.DB.QueryRow("SELECT nome FROM dim_competencia WHERE id = (select max(id) from dim_competencia)")

	if err := row.Scan(&competence.Nome); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Competence{}.Nome, nil
		}
		return Competence{}.Nome, err
	}
	return competence.Nome, nil
}
