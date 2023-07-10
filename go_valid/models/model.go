package models

type Form struct{
	Firstname string `json:"firstname,omitempty"`
	Lastname string `json:"lastname,omitempty"`
	Email string `json:"email,omitempty"`
	Phonenumber int64 `json:"phonenumber,omitempty"`
	City string `json:"city,omitempty"`
}