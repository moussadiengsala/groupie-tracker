package utils

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"learn.zone01dakar.sn/groupie-tracker/core/models"
)

func Fetcher(url string) ([]byte, error) {
	response, serverError := http.Get(url)
	if serverError != nil {
		return nil, serverError
	}
	defer response.Body.Close()
	return ioutil.ReadAll(response.Body)
}

func ConverToInt(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}
	return num
}

func ConvertArrayToInt(str []string) []int {
	var res []int
	for _, v := range str {
		res = append(res, ConverToInt(v))
	}
	return res
}

func ConvertArrayToString(str []int) []string {
	var res []string
	for _, v := range str {
		res = append(res, strconv.Itoa(v))
	}
	return res
}

func IsContainsOnArray(array []string, str string) bool {
	for _, value := range array {
		if strings.Contains(value, str) {
			return true
		}
	}
	return false
}

func FilterForGettingUniqueLocation(data []models.SingleArtist) []string {
	var output []string
	// fullArtistsInfo, err := Controllers.GetFullArtistDetails()
	for _, value := range data {
		for _, addr := range value.Location.Locations {
			if !IsContainsOnArray(output, addr) {
				output = append(output, addr)
			}
		}
	}
	return output
}

func Filter(DataToFilter []models.SingleArtist, filters models.Filters) ([]models.SingleArtist, error) {

	var FilteredDatas = []models.SingleArtist{}
	var isNumberOfMembersMatched = func(value models.SingleArtist) bool {
		if len(filters.NumberOfMembers) == 0 {
			return true
		}
		for _, v := range filters.NumberOfMembers {
			if len(value.Artist.Members) == v {
				return true
			}
		}
		return false
	}

	var isInCreationDateRange = func(value models.SingleArtist) bool {
		return value.Artist.CreationDate >= filters.CreationDate[0] && value.Artist.CreationDate <= filters.CreationDate[1]
	}

	var isInFirstAlbumYearRange = func(value models.SingleArtist) bool {
		var firstAlbum = strings.Split(value.Artist.FirstAlbum, "-")
		var firstAlbumYear, _ = strconv.Atoi(firstAlbum[len(firstAlbum)-1])
		return firstAlbumYear >= filters.FirstAlbumRelease[0] && firstAlbumYear <= filters.FirstAlbumRelease[1]
	}

	var f = func(str ...string) bool {
		for _, v := range str {
			if strings.Contains(strings.ToLower(v), strings.ToLower(filters.Shearch)) {
				return true
			}
		}
		return false
	}

	for _, data := range DataToFilter {
		if filters.Shearch == "" {
			if isNumberOfMembersMatched(data) && IsContainsOnArray(data.Location.Locations, filters.Location) && isInCreationDateRange(data) && isInFirstAlbumYearRange(data) {
				FilteredDatas = append(FilteredDatas, data)
			}
		} else {
			if f(data.Artist.Name, strings.Join(data.Artist.Members, " "), data.Artist.FirstAlbum, data.Artist.ConcertDates) {

				FilteredDatas = append(FilteredDatas, data)
			}
		}
	}

	return FilteredDatas, nil
}

func ErrorThrower(w http.ResponseWriter, errorCode int, message string) {
	tmpl := template.Must(template.ParseFiles("static/templates/error.html"))
	w.WriteHeader(errorCode)
	tmpl.Execute(w, models.Error{ErrorCode: errorCode, ErrorMessage: message})
}

func Parse[T models.FetchedType](data []byte, d T) (T, error) {
	err := json.Unmarshal(data, &d)
	if err != nil {
		return d, err
	}
	return d, nil
}
