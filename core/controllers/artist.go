package Controllers

import (
	"encoding/json"
	"errors"
	"sync"

	"learn.zone01dakar.sn/groupie-tracker/core/models"
	"learn.zone01dakar.sn/groupie-tracker/core/utils"
)

func ParseDataArtists(data []byte) ([]models.Artist, error) {
	var artists []models.Artist
	err := json.Unmarshal(data, &artists)
	if err != nil {
		return nil, err
	}
	return artists, nil
}

/*
This file implements all the methods for an artist
*/
func GetSingleArtist(id string) (models.SingleArtist, error) {
	var artist models.Artist
	var relation models.Relation
	var dates models.Date

	var localFetch = func(str string) ([]byte, error) {
		var resp, err = utils.Fetcher("https://groupietrackers.herokuapp.com/api/" + str + "/" + id)
		if err != nil {
			return nil, err
		}
		return resp, nil
	}
	var responseForArtist, erra = localFetch("artists")
	var responseForRelation, errr = localFetch("relation")
	var responseForDate, errd = localFetch("dates")
	if errr != nil || erra != nil || errd != nil {
		return models.SingleArtist{}, erra
	}
	jsonErrorA := json.Unmarshal(responseForArtist, &artist)
	if jsonErrorA != nil {
		return models.SingleArtist{}, jsonErrorA
	}
	jsonErroR := json.Unmarshal(responseForRelation, &relation)
	if jsonErroR != nil {
		return models.SingleArtist{}, jsonErroR
	}
	jsonErroD := json.Unmarshal(responseForDate, &dates)
	if jsonErroD != nil {
		return models.SingleArtist{}, jsonErroD
	}
	return models.SingleArtist{Artist: artist, Relation: relation, Dates: GetImportantDates(dates)}, nil
}

func GetAllArtists() ([]models.Artist, error) {

	resp, err := utils.Fetcher("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		return nil, err
	}

	artists, errParse := ParseDataArtists(resp)
	if errParse != nil {
		return nil, errParse
	}

	return artists, nil
}

type Date struct {
	Artist   []models.Artist
	Relation []models.Relation
	Location []models.Locations
}

func getAllItems() (Date, error) {
	var wg sync.WaitGroup
	var artists = make(chan []models.Artist)
	var relation = make(chan []models.Relation)
	var location = make(chan []models.Locations)

	var err = make(chan error)
	done := make(chan struct{})

	wg.Add(1)
	go func() {
		defer wg.Done()
		var loc, errl = GetAllLocations()
		if errl != nil {
			err <- errl
			return
		}
		select {
		case location <- loc:
		case <-done:
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		var artist, erra = GetAllArtists()
		if erra != nil {
			err <- erra
			return
		}
		select {
		case artists <- artist:
		case <-done:
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		var rel, errr = GetAllRelations()
		if errr != nil {
			err <- errr
			return
		}
		select {
		case relation <- rel:
		case <-done:
		}
	}()

	go func() {
		wg.Wait()
		close(artists)
		close(relation)
		close(location)
		close(err)
	}()

	select {
	case artists := <-artists:
		relation := <-relation
		location := <-location
		return Date{artists, relation, location}, nil
	case err := <-err:
		return Date{}, err
	case <-done:
		return Date{}, errors.New("operation canceled")
	}
}

func GetFullArtistDetails() ([]models.SingleArtist, error) {
	var singleArtist []models.SingleArtist
	var dates, err = getAllItems()

	if err != nil {
		return nil, err
	}

	for idx, artist := range dates.Artist {
		singleArtist = append(singleArtist, models.SingleArtist{Artist: artist, Relation: dates.Relation[idx], Location: dates.Location[idx]})
	}
	return singleArtist, err
}
