package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"maps"

	"github.com/julienschmidt/httprouter"
)

func (app *application) readIDParams(r *http.Request) (int64, error) {
	params := httprouter.ParamsFromContext(r.Context())
	idstr := params.ByName("id")
	id, err := strconv.ParseInt(idstr, 10, 64)
	if err != nil || id < 1 {
		return 0, errors.New("invalid params id")
	}
	return id, err
}

func (app *application) writeJSON(w http.ResponseWriter, status int, data any, headers http.Header) error {
	// Create envelope for consistent response format
	envelope := map[string]any{
		"success": true,
		"data":    data,
	}

	js, err := json.Marshal(envelope)

	if err != nil {
		return err
	}
	js = append(js, '\n')

	maps.Copy(w.Header(), headers)
	w.Header().Set("Content-Type",
		"application/json")
	w.WriteHeader(status)
	w.Write(js)
	return nil
}

// errorJSON sends a JSON-formatted error response with the provided error message and status code
func (app *application) errorJSON(w http.ResponseWriter, err error, statusCode int) {
	errorMessage := "The server encountered a problem and could not process your request"
	if err != nil {
		errorMessage = err.Error()
	}

	// Create response envelope with success=false and null data
	envelope := map[string]any{
		"success": false,
		"data":    nil,
		"msg":     errorMessage,
	}

	js, err := json.Marshal(envelope)
	if err != nil {
		app.logger.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	js = append(js, '\n')

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(js)
}
