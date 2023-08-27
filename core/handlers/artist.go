package handlers

import (
	// "fmt"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"text/template"

	Controllers "learn.zone01dakar.sn/groupie-tracker/core/controllers"
	"learn.zone01dakar.sn/groupie-tracker/core/models"
	"learn.zone01dakar.sn/groupie-tracker/core/utils"
)

func AllArtist(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var fullArtistsInfo, err = Controllers.GetFullArtistDetails()
	// var AllLocations = utils.FilterForGettingUniqueLocation(fullArtistsInfo)

	if err != nil {
		utils.ErrorThrower(w, http.StatusInternalServerError, "Internal error")
		return
	}
	filters := models.Filters{
		NumberOfMembers:   utils.ConvertArrayToInt((r.Form["members"])),
		CreationDate:      []int{utils.ConverToInt(r.FormValue("creation-date-0")), utils.ConverToInt(r.FormValue("creation-date-1"))},
		FirstAlbumRelease: []int{utils.ConverToInt(r.FormValue("first-album-release-0")), utils.ConverToInt(r.FormValue("first-album-release-1"))},
		Location:          r.FormValue("location"),
	}

	if r.Method == http.MethodPost {
		datas, errfilter := utils.Filter(fullArtistsInfo, filters)
		fullArtistsInfo = datas
		if errfilter != nil {
			utils.ErrorThrower(w, http.StatusInternalServerError, "Internal error")
			return
		}

	}
	tmpl, errparse := template.ParseFiles("static/templates/artists.html")
	if errparse != nil {
		utils.ErrorThrower(w, http.StatusInternalServerError, "Internal error")
		fmt.Println(errparse)
		return
	}
	var forLocationPurpose, _ = Controllers.GetFullArtistDetails()

	tmpl.Execute(w, map[string]interface{}{
		"Datas": fullArtistsInfo,
		// "Locations":         AllLocations,
		"Members":           []int{1, 2, 3, 4, 5, 6, 7,8},
		"CreateDate":        []string{"1940", "2023"},
		"FirtAlbumReleases": []string{"1940", "2023"},
		"LocationConcert":   utils.FilterForGettingUniqueLocation(forLocationPurpose),
	})
	fullArtistsInfo, err = Controllers.GetFullArtistDetails()
}

func SingleArtist(w http.ResponseWriter, r *http.Request) {
	splittedURL := strings.Split(r.URL.Path, "/")

	if r.Method != "GET" {
		utils.ErrorThrower(w, http.StatusMethodNotAllowed, "Not allowed")
		return
	}

	if len(splittedURL) != 3 || splittedURL[1] != "artists" {
		utils.ErrorThrower(w, http.StatusNotFound, "Not Found")
		return
	}
	artistId := splittedURL[len(splittedURL)-1]
	var datas, apiError = Controllers.GetSingleArtist(artistId)
	if apiError != nil {
		utils.ErrorThrower(w, http.StatusInternalServerError, "Something went wrong")
		return
	}
	if datas.Artist.Id == 0 {
		utils.ErrorThrower(w, http.StatusNotFound, "Not Found")
		return
	}
	temp, errparse := template.ParseFiles("static/templates/artist.html")
	if errparse != nil {
		utils.ErrorThrower(w, http.StatusInternalServerError, "Internal error")
		return
	}
	temp.Execute(w, map[string]interface{}{
		"Id":    artistId,
		"Datas": datas,
	})
}

func Map(w http.ResponseWriter, r *http.Request) {
	var file, _ = os.ReadFile("mapApi.json")

	if r.Method != "GET" {
		utils.ErrorThrower(w, http.StatusMethodNotAllowed, "Not allowed")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(file)
}

func MapInfoArtists(w http.ResponseWriter, r *http.Request) {
	var fullArtistsInfo, err = Controllers.GetFullArtistDetails()

	if r.Method != "GET" {
		utils.ErrorThrower(w, http.StatusMethodNotAllowed, "Not allowed")
		return
	}
	if err != nil {
		utils.ErrorThrower(w, http.StatusInternalServerError, "Internal error")
		return
	}
	var jsonData, _ = json.MarshalIndent(fullArtistsInfo, "", "   ")
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func AllDataArtist(w http.ResponseWriter, r *http.Request) {
	var fullArtistsInfo, err = Controllers.GetFullArtistDetails()
	if r.Method != "GET" {
		utils.ErrorThrower(w, http.StatusMethodNotAllowed, "Not allowed")
		return
	}
	if err != nil {
		utils.ErrorThrower(w, http.StatusInternalServerError, "Internal error")
		return
	}
	var jsonData, _ = json.MarshalIndent(fullArtistsInfo, "", "\t")
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func SingleDataArtist(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		utils.ErrorThrower(w, http.StatusMethodNotAllowed, "Not allowed")
		return
	}
	splittedURL := strings.Split(r.URL.Path, "/")
	if len(splittedURL) != 3 {
		utils.ErrorThrower(w, http.StatusNotFound, "Not Found")
		return
	}
	artistId := splittedURL[len(splittedURL)-1]
	var singleArtist, err = Controllers.GetSingleArtist(artistId)
	if err != nil {
		utils.ErrorThrower(w, http.StatusInternalServerError, "Internal error")
		return
	}
	var jsonData, _ = json.MarshalIndent(singleArtist, "", "\t")
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
