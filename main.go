package main

import (
	"net/http"

	"learn.zone01dakar.sn/groupie-tracker/core/handlers"
	"learn.zone01dakar.sn/groupie-tracker/core/server"
	// "learn.zone01dakar.sn/groupie-tracker/packages/server"
)

func main() {
	// api.Map()
	// This where to declare routes and their handlers
	routes := map[string]func(http.ResponseWriter, *http.Request){
		"/":              handlers.Root,
		"/artists":       handlers.AllArtist,
		"/artists/":      handlers.SingleArtist,
		"/map":           handlers.Map,
		"/allartists":    handlers.AllDataArtist,
		"/singleartist/": handlers.SingleDataArtist,
	}
	// serve the files under the "/static/" URL path
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	for route, handler := range routes {
		http.HandleFunc(route, handler)
	}
	server.Config(":8000")
}
