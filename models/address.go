package models

type Address struct {
	ID              string `json:"id"`
	Type            string `json:"type"`
	City            string `json:"city"`
	District        string `json:"district"`
	Neighborhood    string `json:"neighborhood"`
	Description     string `json:"description"`
	PhoneNumber     string `json:"phone_number"`
	Floor           string `json:"floor"`
	DoorNumber      string `json:"door_number"`
	ApartmentNumber string `json:"apartment_number"`
	PostalCode      string `json:"postal_code"`
}
