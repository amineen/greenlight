package main

import (
	"encoding/json"
	"net/http"
)

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"status":      "available",
		"environment": app.config.env,
		"version":     version,
	}
	err := app.writeJSON(w, http.StatusOK, data, nil)
	if err != nil {
		app.logger.Println(err)
		http.Error(w,
			"The server encountered a problem and could not process your request", http.StatusInternalServerError)
	}
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
