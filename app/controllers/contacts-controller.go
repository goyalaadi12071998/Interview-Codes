package controllers

import (
	"encoding/json"
	"interview/app/contacts"
	errorclass "interview/app/error"
	"interview/app/structs"
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

func (c contactcontroller) Identify(w http.ResponseWriter, r *http.Request) {
	identifyData := new(structs.RequestIdentify)
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(identifyData)
	if err != nil {
		Respond(w, r, nil, errorclass.NewError(errorclass.BadRequestError).Wrap("error in data format"))
		return
	}

	err = validateIdentifyRequestData(identifyData)
	if err != nil {
		Respond(w, r, nil, errorclass.NewError(errorclass.BadRequestValidationError).Wrap(err.Error()))
		return
	}

	payload, errr := c.service.CreateContacts(r.Context(), identifyData)
	if errr != nil {
		Respond(w, r, nil, errr)
		return
	}

	Respond(w, r, payload, nil)
}
