package contacts

import (
	"context"
	"interview/app/constants"
	errorclass "interview/app/error"
	"interview/app/models"
	"interview/app/structs"
)

var Service service

type service struct {
	core ICore
}

func NewService(core ICore) IService {
	Service = service{
		core: core,
	}
	return Service
}

func (s service) CreateContact(ctx context.Context, data *structs.RequestIdentify) (*structs.ResponseIdentify, *errorclass.Error) {
	var contactEmail *models.Contact
	var contactPhoneNumber *models.Contact
	var filter map[string]interface{}
	var err error
	var primaryContactId int
	var secondaryContactId int

	if data.Email != "" {
		filter = map[string]interface{}{
			"email": data.Email,
		}
		contactEmail, err = s.core.GetContact(ctx, filter)
		if err != nil {
			return nil, errorclass.NewError(errorclass.InternalServerError).Wrap(err.Error())
		}
	}

	if data.PhoneNumber != "" {
		filter = map[string]interface{}{
			"phone_number": data.PhoneNumber,
		}
		contactPhoneNumber, err = s.core.GetContact(ctx, filter)
		if err != nil {
			return nil, errorclass.NewError(errorclass.InternalServerError).Wrap("internal server error")
		}
	}

	if contactEmail != nil && contactPhoneNumber != nil && contactEmail.Id == contactPhoneNumber.Id {

		if contactEmail.LinkPreference == constants.LinkPreferencePrimary {
			primaryContactId = contactEmail.Id
			secondaryContactId = contactPhoneNumber.Id
		} else {
			primaryContactId = contactPhoneNumber.Id
			secondaryContactId = contactEmail.Id
		}

		return &structs.ResponseIdentify{
			Contact: structs.Contact{
				PrimaryContatctId:   primaryContactId,
				Emails:              []string{contactEmail.Email, contactPhoneNumber.Email},
				PhoneNumbers:        []string{contactEmail.PhoneNumber, contactPhoneNumber.PhoneNumber},
				SecondaryContactIds: []int{secondaryContactId},
			},
		}, nil
	}

	newContact := models.Contact{
		PhoneNumber: data.PhoneNumber,
		Email:       data.Email,
	}

	if contactEmail == nil && contactPhoneNumber == nil {
		newContact.LinkPreference = constants.LinkPreferencePrimary
	} else if contactEmail != nil && contactPhoneNumber == nil {
		newContact.LinkPreference = constants.LinkPreferenceSecondary
		newContact.LinkedId = contactEmail.Id
	} else if contactEmail == nil && contactPhoneNumber != nil {
		newContact.LinkPreference = constants.LinkPreferenceSecondary
		newContact.LinkedId = contactPhoneNumber.Id
	}

	res, err := s.core.CreateContact(ctx, &newContact)
	if err != nil {
		return nil, errorclass.NewError(errorclass.InternalServerError).Wrap("error in saving data in db")
	}

	return &structs.ResponseIdentify{
		Contact: structs.Contact{
			PrimaryContatctId:   res.Id,
			Emails:              []string{res.Email},
			PhoneNumbers:        []string{res.PhoneNumber},
			SecondaryContactIds: []int{},
		},
	}, nil
}
