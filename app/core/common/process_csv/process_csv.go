package process_csv

import (
	"bufio"
	"fmt"
	"main.go/app/core/common/process_csv/transform"
	"main.go/app/core/constants"
	"main.go/app/core/utils/competence"
	"os"
	"strings"
	"time"
)

func ProcessCsv(fileName string, mapper []string, excludeColumns []string) {
	compt := competence.GetCompetence(time.Now())
	csvFileName := constants.ExtractDir + "/" + fileName + compt + ".csv"

	file, err := os.Open(csvFileName)
	if err != nil {
		fmt.Println("Erro ao abrir arquivo CSV!", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var headers []string
	var columnIndex []int

	if scanner.Scan() {
		headers = strings.Split(scanner.Text(), ";")
		for i := range headers {
			headers[i] = strings.ReplaceAll(headers[i], "\"", "")
		}

		for i, header := range headers {
			exclude := false
			for _, excludeColumn := range excludeColumns {
				if header == excludeColumn {
					exclude = true
					break
				}
			}
			if !exclude {
				columnIndex = append(columnIndex, i)
			}
		}
	}

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.ReplaceAll(line, "\"", "")
		line = strings.ReplaceAll(line, ";", ",")
		items := strings.Split(line, ",")
		record := make(map[string]string)

		for i, newName := range mapper {
			if i < len(columnIndex) && columnIndex[i] < len(items) {
				item := items[columnIndex[i]]
				if len(item) == 0 {
					item = "null"
				}
				record[newName] = item
			}
		}

		transform.Transform(record)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Erro ao processar o arquivo CSV", err)
	}
}
