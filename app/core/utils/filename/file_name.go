package filename

import (
	"main.go/app/core/utils/competence"
	"time"
)

func GetFileName(baseFileName string) string {
	var compt = competence.GetCompetence(time.Now())

	baseFileName = baseFileName + compt + ".ZIP"

	return baseFileName
}
