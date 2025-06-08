package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

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

func (app *application) writeJSON(w http.ResponseWriter, status int, data interface{}, headers http.Header) error {
	js, err := json.Marshal(data)

	if err != nil {
		return err
	}
	js = append(js, '\n')

	for key, value := range headers {
		w.Header()[key] = value
	}
	w.Header().Set("Content-Type",
		"application/json")
	w.WriteHeader(status)
	w.Write(js)
	return nil
}
