package server

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"text/template"
)

var temp *template.Template

// This function implements the server configuration
func Config(LOCAL_SERVER_PORT string) {
	// message to show on success
	fmt.Println("Server started at http://localhost" + LOCAL_SERVER_PORT)
	err := http.ListenAndServe(LOCAL_SERVER_PORT, nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		log.Fatalf("error starting server: %s\n", err)
	}
}
