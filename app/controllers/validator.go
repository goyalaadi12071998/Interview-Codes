package controllers

import (
	"errors"
	"interview/app/structs"
)

func validateIdentifyRequestData(data *structs.RequestIdentify) error {
	if data.Email == "" && data.PhoneNumber == "" {
		return errors.New("invalid data")
	}

	return nil
}
