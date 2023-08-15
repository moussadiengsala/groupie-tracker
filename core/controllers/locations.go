package Controllers

import (
	"learn.zone01dakar.sn/groupie-tracker/core/models"
	"learn.zone01dakar.sn/groupie-tracker/core/utils"
)


func GetAllLocations() ([]models.Locations, error) {
	resp, err := utils.Fetcher("https://groupietrackers.herokuapp.com/api/locations")
	if err != nil {
		return nil, err
	}
	locations, errParse := utils.Parse(resp, utils.LocationType{})
	if errParse != nil {
		return nil, errParse
	}
	
	return locations.Index, nil
}

