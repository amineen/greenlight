package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	// Create a map to hold our response data
	data := `{"status":"available", "environment":%q, "version":%q}`

	js := fmt.Sprintf(data, app.config.env, version)

	// Set the content type header to application/json
	w.Header().Set("Content-Type", "application/json")

	w.Write([]byte(js))
}

func (app *application) createMovie(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{
		"msg": "Created a new movie",
	}

	js, err := json.Marshal((response))

	if err != nil {
		app.logger.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {

	id, err := app.readIDParams(r)

	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	response := map[string]int64{
		"id": id,
	}
	js, err := json.Marshal(response)

	if err != nil {
		app.logger.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)

}
