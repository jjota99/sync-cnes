package service

import (
	"log"
	"main.go/app/config/database"
	"main.go/app/modules/competencia/store"
)

type CompetenceService struct {
	CompetenceStore store.CompetenceStore
}

func (s *CompetenceService) GetMaxCompetence() string {
	db, err := database.ConnectDb()

	competenceStore := &store.SQLCompetenceStore{DB: db}

	competence, err := competenceStore.GetMaxCompetence()
	if err != nil {
		log.Fatalf("Falha eu buscar competencia máxima!", err)
	}

	return competence
}
