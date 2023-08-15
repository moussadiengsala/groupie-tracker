package handlers

import (
	"net/http"
	"text/template"

	Controllers "learn.zone01dakar.sn/groupie-tracker/core/controllers"
	"learn.zone01dakar.sn/groupie-tracker/core/utils"
)

func Root(w http.ResponseWriter, r *http.Request) {
	var fullArtistsInfo, err = Controllers.GetFullArtistDetails()
	if err != nil {
		utils.ErrorThrower(w, http.StatusInternalServerError, "Internal error")
		return
	}
	if r.URL.Path != "/" {
		utils.ErrorThrower(w, http.StatusNotFound, "Not Found")
		return
	}
	if r.Method != "GET" {
		utils.ErrorThrower(w, http.StatusMethodNotAllowed, "Not allowed")
		return
	}
	tmpl, errparse := template.ParseFiles("static/templates/index.html")
	if errparse != nil {
		utils.ErrorThrower(w, http.StatusInternalServerError, "Internal error")
		return
	}
	tmpl.Execute(w, map[string]interface{}{
		"Datas": fullArtistsInfo,
	})
}
