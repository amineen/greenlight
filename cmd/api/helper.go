package main

import (
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
