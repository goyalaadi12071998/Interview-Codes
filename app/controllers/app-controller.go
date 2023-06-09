package controllers

import (
	"encoding/json"
	"net/http"
)

var AppController appcontroller

var isHealthy bool = true

type appcontroller struct {
}

func (a appcontroller) Health(w http.ResponseWriter, r *http.Request) {
	if isHealthy {
		w.WriteHeader(200)
		response := map[string]string{
			"message": "Server is healthy",
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(400)
	response := map[string]string{
		"message": "Server is unhealthy",
	}
	json.NewEncoder(w).Encode(response)
	return
}

func (a appcontroller) Get(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	json.NewEncoder(w).Encode("Welcome To Bitespeed")
}

func (a appcontroller) NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{
		"message": "Sorry, this is not a valid endpoint pls check the docs again",
	}

	w.WriteHeader(404)
	json.NewEncoder(w).Encode(response)
}
