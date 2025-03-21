package main

import "time"

type Phone struct {
	TypeID      int `json:"type_id,omitempty"`
	CountryCode int `json:"country_code"`
	Operator    int `json:"operator"`
	Number      int `json:"number"`
}

type Contact struct {
	ID         int       `json:"id"`
	Username   string    `json:"username" validate:"required,ascii"`
	GivenName  string    `json:"given_name,omitempty" validate:"omitempty,alphaunicode"`
	FamilyName string    `json:"family_name,omitempty" validate:"omitempty,alphaunicode"`
	Phones     []Phone   `json:"phones"`
	Email      []string  `json:"email"`
	Birthdate  time.Time `json:"birthdate"`
}

type Group struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Contacts    []int  `json:"contacts"`
}
