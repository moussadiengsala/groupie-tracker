package Controllers

import (
	"learn.zone01dakar.sn/groupie-tracker/core/models"
)

func  GetImportantDates(date models.Date) models.Date{
	var importants []string
	for _, d := range date.Dates {
		if len(d) > 0 && d[0] == '*' {
			importants = append(importants, d[1:])
		}
	}
	date.Dates = importants
	return date
}