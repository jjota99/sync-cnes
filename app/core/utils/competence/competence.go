package competence

import (
	"encoding/json"
	"strconv"
	"strings"
	"time"
)

func GetCompetence(currentDate time.Time) string {
	year := currentDate.Year()
	month := int(currentDate.Month())

	if month == 1 {
		year = year - 1
		month = 12
	}

	month = month - 2
	var convertedYear = strconv.Itoa(year)
	var convertedMonth = strconv.Itoa(month)

	if len(convertedMonth) < 2 {
		convertedMonth = "0" + convertedMonth
	}

	var result, _ = json.Marshal(convertedYear + convertedMonth)

	return strings.Trim(string(result), "\"")
}
