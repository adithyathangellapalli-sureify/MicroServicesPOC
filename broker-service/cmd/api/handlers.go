package main

import (
	"net/http"
)

type JSONPayload struct {
	Data string `json:"data"`
	Name string `json:"name"`
}

func (app *Config) Broker(w http.ResponseWriter, r *http.Request) {

	var requestPayload JSONPayload
	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: "Hitted the broker microservice. Hi " + requestPayload.Name + " This is response we are sending from broker",
	}
	_ = app.writeJSON(w, http.StatusOK, payload)

}
