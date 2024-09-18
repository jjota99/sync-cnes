package listfiles

import (
	"main.go/app/core/utils/competence"
	"time"
)

func GetListFiles(list []string) []string {
	var ListFiles = list
	var competence = competence.GetCompetence(time.Now())

	for i, f := range ListFiles {
		ListFiles[i] = f + competence + ".process_csv"
	}

	return ListFiles
}
