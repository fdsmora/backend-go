package main

import (
	"net/http"
)

func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{"status":"available","environment":"development","version":"1.0.0"}`))
}
