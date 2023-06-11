package structs

type RequestIdentify struct {
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
}

type ResponseIdentify struct {
	Contact Contact `json:"contact"`
}

type Contact struct {
	PrimaryContatctId   int      `json:"primaryContatctId"`
	Emails              []string `json:"emails"`
	PhoneNumbers        []string `json:"phoneNumbers"`
	SecondaryContactIds []int    `json:"secondaryContactIds"`
}
