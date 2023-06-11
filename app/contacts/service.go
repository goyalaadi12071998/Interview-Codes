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

func (s service) CreateContacts(ctx context.Context, data *structs.RequestIdentify) (*structs.ResponseIdentify, *errorclass.Error) {
	var filter map[string]interface{}

	filter = map[string]interface{}{
		"email":        data.Email,
		"phone_number": data.PhoneNumber,
	}

	existingContacts, err := s.core.GetContact(ctx, filter)
	if err != nil {
		return nil, errorclass.NewError(errorclass.InternalServerError).Wrap(err.Error())
	}

	if existingContacts != nil {
		contactData := existingContacts[0]
		findPrimaryContact := false

		var primary models.Contact

		if contactData.LinkPreference != constants.LinkPreferencePrimary {
			findPrimaryContact = true
		}

		if findPrimaryContact == true {
			filter = map[string]interface{}{
				"link_preference": constants.LinkPreferencePrimary,
				"id":              contactData.LinkedId,
			}
			allprimaryContacts, err := s.core.GetContact(ctx, filter)
			if err != nil {
				return nil, errorclass.NewError(errorclass.InternalServerError).Wrap(err.Error())
			}
			primary = allprimaryContacts[0]
		} else {
			primary = contactData
		}

		return s.FetchSecondaryAndReturnIdentifyResponse(ctx, primary)
	}

	var emailContacts []models.Contact
	var phoneContacts []models.Contact

	filter = map[string]interface{}{
		"email": data.Email,
	}
	emailContacts, err = s.core.GetContact(ctx, filter)
	if err != nil {
		return nil, errorclass.NewError(errorclass.InternalServerError).Wrap(err.Error())
	}

	filter = map[string]interface{}{
		"phone_number": data.PhoneNumber,
	}
	phoneContacts, err = s.core.GetContact(ctx, filter)
	if err != nil {
		return nil, errorclass.NewError(errorclass.InternalServerError).Wrap(err.Error())
	}

	if emailContacts != nil && phoneContacts != nil {
		primaryContactEmail := findPrimaryContact(emailContacts)
		primaryContactPhone := findPrimaryContact(phoneContacts)

		if primaryContactEmail.CreatedAt < primaryContactPhone.CreatedAt {
			filter = map[string]interface{}{
				"linked_id":       primaryContactEmail.Id,
				"link_preference": constants.LinkPreferenceSecondary,
			}
			return s.UpdateAndFetchSecondaryAndReturnIdentifyResponse(ctx, primaryContactEmail, primaryContactPhone, filter)
		}

		filter = map[string]interface{}{
			"linked_id":       primaryContactPhone.Id,
			"link_preference": constants.LinkPreferenceSecondary,
		}
		return s.UpdateAndFetchSecondaryAndReturnIdentifyResponse(ctx, primaryContactPhone, primaryContactEmail, filter)
	} else if emailContacts == nil && phoneContacts == nil {
		newcontact := models.Contact{
			PhoneNumber:    data.PhoneNumber,
			Email:          data.Email,
			LinkPreference: constants.LinkPreferencePrimary,
		}

		res, err := s.core.CreateContact(ctx, &newcontact)
		if err != nil {
			return nil, errorclass.NewError(errorclass.InternalServerError).Wrap(err.Error())
		}
		otherContacts := []models.Contact{}
		return formatIdentifyResponse(*res, otherContacts)
	}

	skipCreateNewContact := false
	if (emailContacts == nil && data.Email == "") || (phoneContacts == nil && data.PhoneNumber == "") {
		skipCreateNewContact = true
	}

	var primaryContact models.Contact
	var secondaryContact models.Contact
	var contacts []models.Contact
	foundPrimaryContact := false

	if emailContacts == nil {
		contacts = phoneContacts
	} else {
		contacts = emailContacts
	}

	for _, contact := range contacts {
		if contact.LinkPreference == constants.LinkPreferencePrimary {
			primaryContact = contact
			foundPrimaryContact = true
		}
	}

	if foundPrimaryContact == false {
		secondaryContact = contacts[0]
		filter = map[string]interface{}{
			"link_preference": constants.LinkPreferencePrimary,
			"id":              secondaryContact.LinkedId,
		}
		contacts, err = s.core.GetContact(ctx, filter)
		if err != nil {
			return nil, errorclass.NewError(errorclass.InternalServerError).Wrap(err.Error())
		}

		primaryContact = contacts[0]
	}

	if skipCreateNewContact == false {
		var newcontact = models.Contact{
			PhoneNumber:    data.PhoneNumber,
			Email:          data.Email,
			LinkPreference: constants.LinkPreferenceSecondary,
			LinkedId:       primaryContact.Id,
		}

		_, errr := s.core.CreateContact(ctx, &newcontact)
		if errr != nil {
			return nil, errorclass.NewError(errorclass.InternalServerError).Wrap(err.Error())
		}
	}

	return s.FetchSecondaryAndReturnIdentifyResponse(ctx, primaryContact)
}

func (s service) UpdateAndFetchSecondaryAndReturnIdentifyResponse(ctx context.Context, primary models.Contact, secondary models.Contact, filter map[string]interface{}) (*structs.ResponseIdentify, *errorclass.Error) {
	_, err := s.core.UpdateContact(ctx, &secondary, filter)
	if err != nil {
		return nil, errorclass.NewError(errorclass.InternalServerError).Wrap(err.Error())
	}

	return s.FetchSecondaryAndReturnIdentifyResponse(ctx, primary)
}

func (s service) FetchSecondaryAndReturnIdentifyResponse(ctx context.Context, primary models.Contact) (*structs.ResponseIdentify, *errorclass.Error) {
	var filter map[string]interface{}

	filter = map[string]interface{}{
		"link_preference": constants.LinkPreferenceSecondary,
		"linked_id":       primary.Id,
	}

	secondary, err := s.core.GetContact(ctx, filter)
	if err != nil {
		return nil, errorclass.NewError(errorclass.InternalServerError).Wrap(err.Error())
	}
	return formatIdentifyResponse(primary, secondary)
}

func formatIdentifyResponse(primary models.Contact, secondary []models.Contact) (*structs.ResponseIdentify, *errorclass.Error) {
	secondryContactEmails := []string{}
	secondryContactPhoneNumbers := []string{}
	secondayContactIds := []int{}

	secondryContactEmails = append(secondryContactEmails, primary.Email)
	secondryContactPhoneNumbers = append(secondryContactPhoneNumbers, primary.PhoneNumber)

	for _, each := range secondary {
		if contains(each.Email, secondryContactEmails) {
			secondryContactEmails = append(secondryContactEmails, each.Email)
		}

		if contains(each.PhoneNumber, secondryContactPhoneNumbers) {
			secondryContactPhoneNumbers = append(secondryContactPhoneNumbers, each.PhoneNumber)
		}

		if containsint(each.Id, secondayContactIds) {
			secondayContactIds = append(secondayContactIds, each.Id)
		}
	}

	return &structs.ResponseIdentify{
		Contact: structs.Contact{
			PrimaryContatctId:   primary.Id,
			Emails:              secondryContactEmails,
			PhoneNumbers:        secondryContactPhoneNumbers,
			SecondaryContactIds: secondayContactIds,
		},
	}, nil
}

func contains(data string, array []string) bool {
	for _, x := range array {
		if x == data {
			return false
		}
	}
	return true
}

func containsint(data int, array []int) bool {
	for _, x := range array {
		if x == data {
			return false
		}
	}
	return true
}

func findPrimaryContact(data []models.Contact) models.Contact {
	for _, contact := range data {
		if contact.LinkPreference == constants.LinkPreferencePrimary {
			return contact
		}
	}

	return models.Contact{}
}
