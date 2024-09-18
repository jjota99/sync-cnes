package decider

import (
	"fmt"
	"main.go/app/core/utils/competence"
	"main.go/app/modules/competencia/service"
	"time"
)

func RunDecider() bool {
	currentCompetence := competence.GetCompetence(time.Now())
	competenceService := &service.CompetenceService{}
	maxCompetence := competenceService.GetMaxCompetence()

	if time.Now().Day() < 16 {
		fmt.Println("Rotina deve ser rodada apenas a partir do dia 16!")
		return false
	}

	if currentCompetence == maxCompetence {
		fmt.Println("Competencia máxima existente ja está na versão mais atualizada!")
		return false
	}

	return true
}
