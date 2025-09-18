package model

import "time"

type Admin struct {
	ID            uint64     `json:"id"`                       // bigint unsigned, auto_increment
	UID           string     `json:"uid"`                      // varchar(191), unique
	CardUID       string     `json:"carduid"`                  // varchar(191), indexed
	UIDCreator    string     `json:"uidCreator"`               // varchar(191), indexed
	Subscriber    string     `json:"subscriber"`               // varchar(191), indexed
	PhotoURL      string     `json:"photo_url"`                // varchar(191)
	Weight        int        `json:"weight"`                   // int(20), default 0
	Name          string     `json:"name"`                     // varchar(191)
	Email         string     `json:"email"`                    // varchar(191)
	Password      string     `json:"password"`                 // varchar(191)
	Ccode         string     `json:"Ccode"`                    // varchar(191), indexed
	Phone         string     `json:"phone"`                    // varchar(191)
	PhoneNumber   string     `json:"PhoneNumber"`              // varchar(191), indexed
	Platform      string     `json:"platform"`                 // varchar(191), indexed
	Status        string     `json:"status"`                   // varchar(191), indexed
	Gender        string     `json:"gender"`                   // varchar(191), indexed
	Age           string     `json:"age"`                      // varchar(191), indexed
	InitCountry   string     `json:"initCountry"`              // varchar(191)
	Country       string     `json:"country"`                  // varchar(191), indexed
	MaritalStatus string     `json:"marital_status"`           // varchar(191), indexed
	RememberToken *string    `json:"remember_token,omitempty"` // nullable
	CreatedAt     *time.Time `json:"created_at,omitempty"`     // nullable timestamp
	UpdatedAt     *time.Time `json:"updated_at,omitempty"`     // nullable timestamp
	UserUpdatedAt *time.Time `json:"userUpdated_at,omitempty"` // nullable timestamp
	PackUid       string     `json:"packUid"`                  // varchar(191), indexed
	PackType      string     `json:"packType"`                 // varchar(191)
	StatusValid   string     `json:"statusValid"`              // varchar(191), indexed
}
