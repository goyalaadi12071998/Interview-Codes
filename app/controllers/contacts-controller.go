package controllers

import (
	"interview/app/contacts"
	errorclass "interview/app/error"
	"net/http"
)

var ContactController contactcontroller

type contactcontroller struct {
	service contacts.IService
}

func NewContactsController(service contacts.IService) {
	ContactController = contactcontroller{
		service: service,
	}
}

func (c contactcontroller) Contact(w http.ResponseWriter, r *http.Request) {
	Respond(w, r, nil, errorclass.NewError(errorclass.InternalServerError).Wrap("Nothig is here"))
}
