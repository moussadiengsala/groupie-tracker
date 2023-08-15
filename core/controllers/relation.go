package Controllers

import (
	"learn.zone01dakar.sn/groupie-tracker/core/models"
	"learn.zone01dakar.sn/groupie-tracker/core/utils"
)

func GetAllRelations() ([]models.Relation, error) {
	resp, err := utils.Fetcher("https://groupietrackers.herokuapp.com/api/relation")
	if err != nil {
		return nil, err
	}

	relations, errParse := utils.Parse(resp, utils.RelationType{})
	if errParse != nil {
		return nil, errParse
	}
	
	return relations.Index, nil
}
