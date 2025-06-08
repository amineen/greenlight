package main

import (
	"errors"
	"net/http"
	"time"

	"github.com/amineen/greenlight/internal/data"
)

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"status":      "available",
		"environment": app.config.env,
		"version":     version,
	}
	err := app.writeJSON(w, http.StatusOK, data, nil)
	if err != nil {
		app.errorJSON(w, err, http.StatusInternalServerError)
	}
}

func (app *application) createMovie(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{
		"msg": "Created a new movie",
	}

	err := app.writeJSON(w, http.StatusCreated, response, nil)
	if err != nil {
		app.errorJSON(w, err, http.StatusInternalServerError)
		return
	}
}

func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {

	id, err := app.readIDParams(r)

	if err != nil || id < 1 {
		app.errorJSON(w, errors.New("movie not found"), http.StatusNotFound)
		return
	}

	movie := data.Movie{
		ID:        id,
		CreatedAt: time.Now(),
		Title:     "Casablanca",
		Runtime:   102,
		Year:      2025,
		Genres: []string{
			"drama",
			"romance",
			"war"},
		Version: 1,
	}

	err = app.writeJSON(w, http.StatusOK, movie, nil)
	if err != nil {
		app.errorJSON(w, err, http.StatusInternalServerError)
		return
	}

}
